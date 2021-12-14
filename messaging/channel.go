package messaging

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
	"net"
	"time"
)

const (
	DefaultWriteIdle = 5 * time.Second
)

type Channel struct {
	name string

	conn net.Conn

	readReady chan struct{}

	closing   chan struct{}
	doClosing atomic.Bool

	pendings map[uint32]*operation

	readMsg chan *readMsg

	sending chan *operation
	sent    chan *operation

	reqHandlers      []MsgHandler
	asyncRspHandlers []MsgHandler

	timeout time.Duration

	writeIdle time.Duration
	pingTimer *time.Timer
	pingStop  chan struct{}

	readIdle time.Duration

	closeListener CloseListener

	readErr chan error // errors from read
}

type CloseListener interface {
	ChannelClose(string)
}

type MsgHandler func(*Channel, v1.Message) error

func NewChannelWithConfig(name string, timeout int, writeIdle time.Duration, conn net.Conn) *Channel {
	c := &Channel{name: name, conn: conn, closing: make(chan struct{}, 1), pendings: make(map[uint32]*operation),
		readMsg: make(chan *readMsg), sending: make(chan *operation), sent: make(chan *operation), readReady: make(chan struct{}),
		timeout: time.Duration(timeout) * time.Millisecond, writeIdle: writeIdle, readErr: make(chan error)}

	go c.run()

	c.pingStop = make(chan struct{})
	c.pingTimer = time.NewTimer(c.writeIdle)
	go c.ping()

	<-c.readReady

	return c
}

func NewChannel(conn net.Conn) *Channel {
	return NewChannelWithConfig(conn.RemoteAddr().String(), 0, DefaultWriteIdle, conn)
}

func (c *Channel) ping() {
	for {
		select {
		case <-c.pingTimer.C:
			ping := v1.NewHeartbeatRequest()
			if err := c.AsyncCall(ping); err != nil {
				logging.Errorf("write ping error %+v", err)
				c.Close()
			}

		case <-c.pingStop:
			c.pingTimer.Stop()
			return
		}
	}
}

func (c *Channel) run() {
	go c.read()

	for {
		select {
		case <-c.closing:
			fmt.Println("run closing")

			if !c.doClosing.Load() {
				c.doClosing.CAS(false, true)

				c.pingStop <- struct{}{}

				fmt.Println("conn close")
				if err := c.conn.Close(); err != nil {
					logging.Warningf("closing error %s", err)
				}

				if c.closeListener != nil {
					c.closeListener.ChannelClose(c.name)
				}
			}
			fmt.Println("run exit")
			return

		case err := <-c.readErr:
			logging.Warningf("read error %s, close channel", err)

			for id, op := range c.pendings {
				delete(c.pendings, id)
				op.err = err
				close(op.rsp)
			}

			c.Close()

		case read := <-c.readMsg:

			fmt.Printf("read message type %s\n", read.msg.Tp.String())

			switch read.msg.Tp {
			case v1.MsgTypeHeartbeatResponse:
				fallthrough
			case v1.MsgTypeResponse:
				pending := c.pendings[read.msg.Id]
				if pending == nil {
					logging.Warningf("read message req id %d not found!", read.msg.Id)
				}
				delete(c.pendings, read.msg.Id)

				if pending.reqTp == v1.MsgTypeRequestOneway || pending.reqTp == v1.MsgTypeHeartbeatRequest {
					if read.err == nil {
						for _, handle := range c.asyncRspHandlers {
							if err := handle(c, *read.msg); err != nil {
								logging.Errorf("handling async")
							}
						}
					}

				} else {
					pending.err = read.err
					pending.rsp <- read.msg
				}

			case v1.MsgTypeRequestSync:
				fallthrough
			case v1.MsgTypeRequestOneway:

				if read.err == nil {
					for _, handle := range c.reqHandlers {
						if err := handle(c, *read.msg); err != nil {
							logging.Errorf("handling request message error %+v", err)
						}
					}
				}

			case v1.MsgTypeHeartbeatRequest:
				fmt.Println("how to handle message heartbeat request")

			default:
				logging.Warningf("message type unknown %d", read.msg.Tp)
			}

		case sent := <-c.sent:
			if sent.err != nil {
				delete(c.pendings, sent.id)
			}

		case sending := <-c.sending:
			c.pendings[sending.id] = sending
		}
	}

}

type readMsg struct {
	msg *v1.Message
	err error
}

func (c *Channel) read() {

	c.readReady <- struct{}{}

	for {

		if c.doClosing.Load() {
			fmt.Println("read exit")
			// todo: handling msg already read ?
			return
		}

		fmt.Printf("read ...\n")

		if c.timeout != 0 {
			c.conn.SetReadDeadline(time.Now().Add(c.timeout))
		}

		msg, err := v1.ReadMessage(c.conn)

		fmt.Printf("after read msg\n")

		if err != nil {
			/*if err == v1.NotSeataMessage || err == v1.MessageFormatError {
				logging.Warning("read message error")
				continue
			} else {
				if err.Error() == "EOF" {
					logging.Debugf("read %s eof exit", c.name)
				} else if strings.Contains(err.Error(), "use of closed network connection") {
					logging.Debugf("remote close %s ", c.name)
				} else {
					logging.Warningf("read error %s, read exit", err)
				}

				// close directly?
				c.Close()
				return
			}*/
			c.readErr <- err
			return
		}

		c.readMsg <- &readMsg{msg: msg, err: err}
	}
}

func (c *Channel) RegisterRequestHandler(h MsgHandler) {
	c.reqHandlers = append(c.reqHandlers, h)
}

func (c *Channel) RegisterAsyncRspHandler(h MsgHandler) {
	c.asyncRspHandlers = append(c.asyncRspHandlers, h)
}

type operation struct {
	id    uint32
	rsp   chan *v1.Message
	err   error
	reqTp v1.MessageType
}

func (op *operation) wait(ctx context.Context) (rsp v1.Message, err error) {
	fmt.Printf("waiting on id %d\n", op.id)

	select {
	case <-ctx.Done():
		return v1.Message{}, ctx.Err()
	case msg := <-op.rsp:
		if msg == nil {
			rsp = v1.Message{}
		} else {
			rsp = *msg
		}
		err = op.err
		return
	}
}

func (c *Channel) Call(req v1.Message) (rsp v1.Message, err error) {
	ctx := context.Background()
	return c.CallWithCtx(ctx, req)
}

// CallWithCtx if ctx set deadline, it will override the timeout of channel
func (c *Channel) CallWithCtx(ctx context.Context, req v1.Message) (rsp v1.Message, err error) {
	bs, err := v1.EncodeMessage(&req)
	if err != nil {
		return
	}

	op := &operation{id: req.Id, rsp: make(chan *v1.Message, 1)}

	logging.Debugf("call %d \n", req.Id)

	if err = c.send(ctx, op, bs); err != nil {
		return
	}

	rsp, err = op.wait(ctx)
	return
}

func (c *Channel) AsyncCall(msg v1.Message) error {
	ctx := context.Background()
	return c.Async(ctx, msg)
}

func (c *Channel) Async(ctx context.Context, msg v1.Message) error {
	data, err := v1.EncodeMessage(&msg)
	if err != nil {
		return err
	}

	op := &operation{id: msg.Id, rsp: make(chan *v1.Message, 1), reqTp: v1.MsgTypeRequestOneway}

	fmt.Printf("async sending %d \n", msg.Id)

	return c.send(ctx, op, data)
}

func (c *Channel) SendResponse(ctx context.Context, msg *v1.Message) error {
	bs, err := v1.EncodeMessage(msg)
	if err != nil {
		return err
	}

	logging.Debugf("send response %d\n", msg.Id)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return c.write(bs, false)
	}
	return nil
}

func (c *Channel) send(ctx context.Context, op *operation, data []byte) error {
	select {
	case c.sending <- op:

		deadline, ok := ctx.Deadline()
		if ok {
			c.conn.SetWriteDeadline(deadline)
		} else {
			if c.timeout != 0 {
				deadline = time.Now().Add(c.timeout)
				c.conn.SetWriteDeadline(deadline)
			}
		}

		err := c.write(data, false)
		op.err = err
		c.sent <- op
		if err != nil {
			return err
		}

		c.pingTimer.Reset(c.writeIdle)

	case <-ctx.Done():
		fmt.Println("send done")
		// cancel?
		err := ctx.Err()
		return err
	}
	return nil
}

func (c *Channel) write(data []byte, retry bool) error {
	fmt.Printf("write....\n")

	n := 0
	for n < len(data) {
		nt, err := c.conn.Write(data)
		if err != nil {
			return errors.WithStack(err)
		}
		n += nt
	}

	fmt.Printf("write finished\n")
	return nil
}

func (c *Channel) Close() {
	c.closing <- struct{}{}
}

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

)

type Channel struct {
	name string

	conn net.Conn

	closing chan struct{}
	close   atomic.Bool

	pendings map[uint32]*operation

	readOp chan *readMsg

	sending chan *operation
	sent    chan *operation

	reqHandlers []MsgHandler
	asyncRspHandlers []MsgHandler

	timeout time.Duration
}

type MsgHandler func(channel *Channel, message v1.Message) error

func NewChannelWithName(name string, conn net.Conn) Channel {
	c := Channel{name: name, conn: conn, closing: make(chan struct{}), close: atomic.Bool{}, pendings: make(map[uint32]*operation), readOp: make(chan *readMsg),
		sending: make(chan *operation), sent: make(chan *operation)}
	go c.run()
	return c
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{name: conn.RemoteAddr().String(), conn: conn, closing: make(chan struct{}), close: atomic.Bool{}, pendings: make(map[uint32]*operation), readOp: make(chan *readMsg),
		sending: make(chan *operation), sent: make(chan *operation)}
	go c.run()
	return c
}

func (c *Channel) run() {
	go c.read()

	for {
		select {
		case <-c.closing:
			fmt.Println("closing")
			c.close.CAS(false, true)
			if err := c.conn.Close(); err != nil {
				logging.Warningf("closing error", err)
			}

			return

		case read := <-c.readOp:

			//
			if read.err != nil {
				logging.Errorf("read error closing")
				c.closing <- struct{}{}
			}

			fmt.Printf("read message type %s\n", read.msg.Tp.String())

			switch read.msg.Tp {
			case v1.MsgTypeResponse:
				pending := c.pendings[read.msg.Id]
				if pending == nil {
					logging.Warningf("read message req id %d not found!", read.msg.Id)
				}
				delete(c.pendings, read.msg.Id)

				if pending.reqTp==v1.MsgTypeRequestOneway {
					for _, handle := range c.asyncRspHandlers {
						if err := handle(c, *read.msg);err!=nil {
							logging.Errorf("handling async")
						}
					}

				}else {
					pending.rsp <- read.msg
					pending.err = read.err
				}

			case v1.MsgTypeRequestSync:
				fallthrough
			case v1.MsgTypeRequestOneway:

				for _, handle := range c.reqHandlers {
					if err := handle(c, *read.msg); err != nil {
						logging.Errorf("handling request message error %+v", err)
					}
				}

			case v1.MsgTypeHeartbeatRequest:

			case v1.MsgTypeHeartbeatResponse:

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
	for {

		fmt.Printf("read ...\n")

		if c.timeout != 0 {
			c.conn.SetReadDeadline(time.Now().Add(c.timeout))
		}

		msg, err := v1.ReadMessage(c.conn)

		fmt.Printf("after read msg\n")

		if c.close.Load() {
			fmt.Println("close, exit read")
			// todo: handle read msg
			return
		}

		read := &readMsg{msg: msg, err: err}
		c.readOp <- read
		if err != nil {
			logging.Warningf("read error %+v, exit read", err)
			break
		}
	}
}

func (c *Channel) RegisterRequestHandler(h MsgHandler) {
	c.reqHandlers = append(c.reqHandlers, h)
}

func (c *Channel) RegisterAsyncRspHandler(h MsgHandler) {
	c.asyncRspHandlers= append(c.asyncRspHandlers, h)
}

type operation struct {
	id  uint32
	rsp chan *v1.Message
	err   error
	reqTp v1.MessageType
}

func (op *operation) wait(ctx context.Context) (rsp v1.Message, err error) {
	fmt.Printf("waiting on id %d\n", op.id)

	select {
	case <-ctx.Done():
		return v1.Message{}, ctx.Err()
	case rsp := <-op.rsp:
		return *rsp, op.err
	}
}

func (c *Channel) Call(req v1.Message) (rsp v1.Message, err error) {
	ctx := context.Background()
	return c.CallWithCtx(ctx, req)
}

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

func (c *Channel) AsyncCall(msg *v1.Message) error {
	ctx := context.Background()
	return c.Async(ctx, msg)
}

func (c *Channel) Async(ctx context.Context, msg *v1.Message) error {
	bs, err := v1.EncodeMessage(msg)
	if err != nil {
		return err
	}

	op := &operation{id: msg.Id, rsp: make(chan *v1.Message, 1), reqTp: v1.MsgTypeRequestOneway}

	fmt.Printf("async sending %d \n", msg.Id)

	return c.send(ctx, op, bs)
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

		// right now use ctx deadline first
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

	case <-ctx.Done():
		// when this happened ?
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
	select {
	case c.closing <- struct{}{}:
	}
}

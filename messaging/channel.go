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
	DefaultWriteIdle       = 5 * time.Second
	DefaultRpcTimeout      = 30 * time.Second
	DefaultEnableHeartbeat = true
)

var (
	channelId atomic.Int64
)

func nextChannelId() int64 {
	return channelId.Inc()
}

type ChannelConfig struct {
	Timeout time.Duration

	WriteIdle       time.Duration
	EnableHeartBeat bool
}

func DefaultConfig() *ChannelConfig {
	return &ChannelConfig{Timeout: DefaultRpcTimeout, EnableHeartBeat: DefaultEnableHeartbeat, WriteIdle: DefaultWriteIdle}
}

type Channel struct {
	config *ChannelConfig

	id string

	conn net.Conn

	readReady chan struct{}

	closing   chan struct{}
	doClosing atomic.Bool

	pendings map[uint32]*operation

	readMsg chan *readMsg

	sending chan *operation
	sent    chan *operation

	syncReqHandlers  []MessageHandler
	asyncReqHandlers []MessageHandler
	asyncRspHandlers []MessageHandler

	pingTimer *time.Timer
	pingStop  chan struct{}

	readIdle time.Duration

	closeListener CloseListener

	readErr chan error // errors from read
}

func (c Channel) String() string {
	return fmt.Sprintf("%s, %s -> %s", c.id, c.conn.LocalAddr().String(), c.conn.RemoteAddr().String())
}

func NewChannelWithConfig(conn net.Conn, config *ChannelConfig) *Channel {
	id := fmt.Sprintf("channel-%d", nextChannelId())
	c := &Channel{id: id, conn: conn, closing: make(chan struct{}, 1), pendings: make(map[uint32]*operation),
		readMsg: make(chan *readMsg), sending: make(chan *operation), sent: make(chan *operation), readReady: make(chan struct{}),
		config: config, readErr: make(chan error)}

	go c.run()

	if c.config.EnableHeartBeat {
		c.pingStop = make(chan struct{})
		c.pingTimer = time.NewTimer(c.config.WriteIdle)
		go c.ping()
	}

	<-c.readReady

	return c
}

func NewChannel(conn net.Conn) *Channel {
	return NewChannelWithConfig(conn, DefaultConfig())
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
			logging.Debugf("channel [%s] closing", c.String())

			if !c.doClosing.Load() {
				c.doClosing.CAS(false, true)

				if c.config.EnableHeartBeat {
					c.pingStop <- struct{}{}
				}

				if err := c.conn.Close(); err != nil {
					logging.Warnf("closing error %s", err)
				}

				if c.closeListener != nil {
					c.closeListener.ChannelClose(c.id)
				}
			}
			return

		case err := <-c.readErr:
			logging.Warnf("read error %s, close channel %s", err, c.String())

			for id, op := range c.pendings {
				delete(c.pendings, id)
				op.err = err
				close(op.rsp)
			}

			c.Close()

		case read := <-c.readMsg:
			switch read.msg.Tp {
			case v1.MsgTypeHeartbeatResponse:
				// todo: heartbeat request/response handler
				fallthrough
			case v1.MsgTypeResponse:
				pending := c.pendings[read.msg.Id]
				if pending == nil {
					logging.Warnf("read message req id %d not found!", read.msg.Id)
				}
				delete(c.pendings, read.msg.Id)

				if pending.reqTp == v1.MsgTypeRequestOneway || pending.reqTp == v1.MsgTypeHeartbeatRequest {
					if read.err == nil {
						for _, handler := range c.asyncRspHandlers {
							if err := handler.HandleMessage(*read.msg); err != nil {
								logging.Errorf("handling async")
							}
						}
					}

				} else {
					pending.err = read.err
					pending.rsp <- read.msg
				}

			case v1.MsgTypeRequestSync:
				if read.err == nil {
					for _, h := range c.syncReqHandlers {
						if err := h.HandleMessage(*read.msg); err != nil {
							logging.Errorf("handling request message error %+v", err)
						}
					}
				}

			case v1.MsgTypeHeartbeatRequest:
				fallthrough
			case v1.MsgTypeRequestOneway:

				if read.err == nil {
					for _, h := range c.asyncReqHandlers {
						if err := h.HandleMessage(*read.msg); err != nil {
							logging.Errorf("handling request message error %+v", err)
						}
					}
				}

			default:
				logging.Warnf("message type unknown %d", read.msg.Tp)
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
			logging.Debugf("channel %s read exit", c.String())
			// todo: handling msg already read ?
			return
		}

		if c.config.Timeout != 0 {
			if err := c.conn.SetReadDeadline(time.Now().Add(c.config.Timeout)); err != nil {
				logging.Warnf("channel id %s set read deadline error %s", c.id, err.Error())
			}
		}

		msg, err := v1.ReadMessage(c.conn)

		logging.Debugf("channel [%s] read message %s", c.String(), msg.String())

		if err != nil {
			c.readErr <- err
			return
		}

		c.readMsg <- &readMsg{msg: msg, err: err}
	}
}

func (c *Channel) RegisterSyncRequestHandler(h MessageHandler) {
	c.syncReqHandlers = append(c.syncReqHandlers, h)
}

func (c *Channel) RegisterAsyncRequestHandler(h MessageHandler) {
	c.asyncReqHandlers = append(c.asyncReqHandlers, h)
}

func (c *Channel) RegisterAsyncResponseHandler(h MessageHandler) {
	c.asyncRspHandlers = append(c.asyncRspHandlers, h)
}

type operation struct {
	id    uint32
	rsp   chan *v1.Message
	err   error
	reqTp v1.MessageType
}

func (op *operation) wait(ctx context.Context) (rsp v1.Message, err error) {
	logging.Debugf("waiting response on message id %d", op.id)

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

	logging.Debugf("channel [%s] sending request [%s]", c.String(), req.String())
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

	logging.Debugf("channel [%s] async sending request [%s]", c.String(), msg.String())
	return c.send(ctx, op, data)
}

func (c *Channel) SendResponse(ctx context.Context, msg *v1.Message) error {
	bs, err := v1.EncodeMessage(msg)
	if err != nil {
		return err
	}

	logging.Debugf("channel [%s] send response [%s]", c.String(), msg.String())
	return c.send(ctx, nil, bs)
}

func (c *Channel) send(ctx context.Context, op *operation, data []byte) error {
	deadline, ok := ctx.Deadline()
	if ok {
		if err := c.conn.SetWriteDeadline(deadline); err != nil {
			logging.Warnf("channel id %s set write deadline error %s", c.id, err.Error())
		}
	} else {
		if c.config.Timeout != 0 {
			deadline = time.Now().Add(c.config.Timeout)
			if err := c.conn.SetWriteDeadline(deadline); err != nil {
				logging.Warnf("channel id %s set write deadline error %s", c.id, err.Error())
			}
		}
	}

	if op == nil {
		return c.write(data, false)
	}

	select {
	case c.sending <- op:
		err := c.write(data, false)
		op.err = err
		c.sent <- op

		if c.pingTimer != nil {
			c.pingTimer.Reset(c.config.WriteIdle)
		}
		return err

	case <-ctx.Done():
		logging.Debug("send done")
		// cancel?
		err := ctx.Err()
		return err
	}
}

func (c *Channel) write(data []byte, retry bool) error {
	n := 0
	for n < len(data) {
		nt, err := c.conn.Write(data)
		if err != nil {
			return errors.WithStack(err)
		}
		n += nt
	}
	return nil
}

func (c *Channel) Close() {
	c.closing <- struct{}{}
}

type MessageHandler interface {
	HandleMessage(msg v1.Message) error
}

type CloseListener interface {
	ChannelClose(string)
}

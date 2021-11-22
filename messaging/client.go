package messaging

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"net"
)

var (
	ErrClientQuit = errors.New("client is closed")
)

type Client struct {
	conn net.Conn

	closing chan struct{}

	pendings map[uint32]*operation

	readOp chan *readMsg

	sending chan *operation
	sent    chan *operation

	asyncHandlers  []func(msg proto.Message) error
}

type readMsg struct {
	id  uint32
	msg proto.Message
	err error
}

func NewClient(svrAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", svrAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := &Client{conn: conn, closing: make(chan struct{}), pendings: make(map[uint32]*operation), readOp: make(chan *readMsg),
		sending: make(chan *operation), sent: make(chan *operation)}
	go c.run()
	return c, nil
}

func (c *Client) run() {

	go c.read()

	for {
		select {
		case <-c.closing:
			fmt.Println("closing")
			if err := c.conn.Close(); err != nil {
				logging.Warningf("closing error", err)
			}
			return

		case read := <-c.readOp:
			pending := c.pendings[read.id]
			if pending == nil {
				logging.Warningf("read message req id %d not found!", read.id)
			}
			delete(c.pendings, read.id)
			pending.rsp <- read.msg
			pending.err = read.err
			if read.err != nil {
				logging.Errorf("read error closing")
				c.closing <- struct{}{}
			}
			if pending.tp==v1.MSGTYPE_RESQUEST_ONEWAY {
				for _, handle := range c.asyncHandlers {
					m := <- pending.rsp
					if err:=handle(m);err!=nil {
						logging.Errorf("handling async message error %+v", err)
					}
				}
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

func (c *Client) read() {
	for {

		fmt.Printf("read ...\n")

		reqId, msg, err := v1.ReadMessage(c.conn)

		fmt.Printf("after read msg\n")

		read := &readMsg{id: reqId, msg: msg, err: err}
		c.readOp <- read
		if err != nil {
			logging.Warningf("read error %+v, exit read", err)
			break
		}
	}
}

type operation struct {
	id  uint32
	rsp chan proto.Message
	err error
	tp v1.MessageType
}

func (op *operation) wait(ctx context.Context) (rsp proto.Message, err error) {
	fmt.Printf("waiting on id %d\n", op.id)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case rsp := <-op.rsp:
		return rsp, op.err
	}
}

func (c *Client) SyncCall(req proto.Message) (rsp proto.Message, err error) {
	ctx := context.Background()
	return c.Call(ctx, req)
}

func (c *Client) Call(ctx context.Context, req proto.Message) (rsp proto.Message, err error) {
	id, bs, err := v1.EncodeMessage(v1.MSGTYPE_RESQUEST_SYNC, req)
	if err != nil {
		return
	}

	op := &operation{id: id, rsp: make(chan proto.Message), tp:v1.MSGTYPE_RESQUEST_SYNC}

	fmt.Printf("sending %d \n", id)

	select {
	case c.sending <- op:
		err = c.write(bs, false)
		op.err = err
		c.sent <- op
		if err != nil {
			return
		}

	case <-ctx.Done():
		err = ctx.Err()
		return
	}

	rsp, err = op.wait(ctx)
	return
}

func (c *Client) AsyncCall(msg proto.Message) error {
	ctx := context.Background()
	return c.Async(ctx, msg)
}

func (c *Client) Async(ctx context.Context, msg proto.Message) error {
	id, bs, err := v1.EncodeMessage(v1.MSGTYPE_RESQUEST_ONEWAY, msg)
	if err != nil {
		return err
	}

	op := &operation{id: id, rsp: make(chan proto.Message, 1), tp:v1.MSGTYPE_RESQUEST_ONEWAY}

	fmt.Printf("async sending %d \n", id)

	select {
	case c.sending <- op:
		err = c.write(bs, false)
		op.err = err
		c.sent <- op
		if err != nil {
			return err
		}

	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

func (c *Client) write(data []byte, retry bool) error {
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

func (c *Client) RegisterAsyncHandler(h func(message proto.Message) error) {
	c.asyncHandlers= append(c.asyncHandlers, h)
}

func (c *Client) Close() {
	select {
	case c.closing <- struct{}{}:
	}
}

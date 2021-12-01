package messaging

import (
	"github.com/pkg/errors"
	"net"
)

var (
	ErrClientQuit = errors.New("client is closed")
)

type Client struct {
	svrAddr string
	*Channel
}

func NewClient(svrAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", svrAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := &Client{svrAddr: svrAddr, Channel: NewChannel(conn)}
	return c, nil
}

func (c *Client) Close() {
	c.Channel.Close()
}

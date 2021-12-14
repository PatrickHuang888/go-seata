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
	return NewClientWithConfig(svrAddr, DefaultConfig())
}

func NewClientWithConfig(svrAddr string, config *ChannelConfig) (*Client, error) {
	conn, err := net.Dial("tcp", svrAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := &Client{svrAddr: svrAddr, Channel: NewChannelWithConfig(conn.RemoteAddr().String(), conn, config)}
	return c, nil
}

func (c *Client) Close() {
	c.Channel.Close()
}

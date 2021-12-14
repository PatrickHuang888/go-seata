package messaging

import (
	"github.com/pkg/errors"
	"net"
	"time"
)

var (
	ErrClientQuit = errors.New("client is closed")
)

type Client struct {
	svrAddr string
	*Channel
}

func NewClient(svrAddr string) (*Client, error) {
	return NewClientWithConfig(svrAddr, 0, DefaultWriteIdle)
}

// NewClientWithConfig timeout is milliseconds
func NewClientWithConfig(svrAddr string, timeout int, writeIdle time.Duration) (*Client, error) {
	conn, err := net.Dial("tcp", svrAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := &Client{svrAddr: svrAddr, Channel: NewChannelWithConfig(conn.RemoteAddr().String(), timeout, writeIdle, conn)}
	return c, nil
}

func (c *Client) Close() {
	c.Channel.Close()
}

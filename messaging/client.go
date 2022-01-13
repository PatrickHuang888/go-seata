package messaging

import (
	"github.com/pkg/errors"
	"net"
)

var (
	ErrClientQuit = errors.New("client is closed")
)

type Client struct {
	appId   string
	txGroup string

	tcAddr string
	*Channel
}

func NewClient(tcAddr string, appId string, txGroup string) (*Client, error) {
	return NewClientWithConfig(tcAddr, appId, txGroup, DefaultConfig())
}

func NewClientWithConfig(tcAddr string, appId string, txGroup string, config *ChannelConfig) (*Client, error) {
	conn, err := net.Dial("tcp", tcAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := &Client{tcAddr: tcAddr, Channel: NewChannelWithConfig(conn, config)}
	return c, nil
}

func (c Client) AppId() string {
	return c.appId
}

func (c Client) TxGroup() string {
	return c.txGroup
}

func (c *Client) Close() {
	c.Channel.Close()
}

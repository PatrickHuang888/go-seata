package messaging

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"go.uber.org/atomic"
	"net"
)

type Server struct {
	addr string

	close        atomic.Bool
	closing      chan struct{}
	closed       chan struct{}
	channelClose chan string

	channels map[string]*Channel

	reqHandlers      []MsgHandler
	asyncRspHandlers []MsgHandler
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, close: atomic.Bool{}, closing: make(chan struct{}, 1), closed: make(chan struct{}, 1),
		channels: map[string]*Channel{}, channelClose: make(chan string)}
}

func (s *Server) Serv() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		logging.Errorf("listen error ", err)
		s.Close()
	}
	logging.Infof("listen on %s\n", s.addr)

	go func() {
		for {
			select {
			case name := <-s.channelClose:
				fmt.Println("server close channel")
				delete(s.channels, name)

			case <-s.closing:
				logging.Info("server closing")

				s.close.CAS(false, true)

				if l != nil {
					if err := l.Close(); err != nil {
						logging.Errorf("listener close error", err)
					}
				}

				for _, ch := range s.channels {
					ch.Close()
				}

				s.closed <- struct{}{}
				return
			}
		}
	}()

loop:
	for {
		c, err := l.Accept()

		if err != nil {
			if s.close.Load() {
				logging.Infof("closing listener")
				break loop
			} else {
				logging.Errorf("accept error %s\n", err.Error())
				continue
			}

		}

		ch := NewChannel(c)
		ch.closeListener = s
		s.channels[ch.name] = ch
		for _, h := range s.reqHandlers {
			ch.RegisterRequestHandler(h)
		}
		for _, h := range s.asyncRspHandlers {
			ch.RegisterAsyncRspHandler(h)
		}

		go ch.run()
	}

	fmt.Println("server closed")
	<-s.closed
}

func (s *Server) RegisterRequestHandler(h MsgHandler) {
	s.reqHandlers = append(s.reqHandlers, h)
}

func (s *Server) RegisterAsyncRspHandler(h MsgHandler) {
	s.asyncRspHandlers = append(s.asyncRspHandlers, h)
}

func (s *Server) Close() {
	s.closing <- struct{}{}
}

func (s *Server) ChannelClose(name string) {
	s.channelClose <- name
}

func handleTmReg(c *Channel, req v1.Message) error {

	_, ok := req.Msg.(*pb.RegisterTMRequestProto)
	if ok {
		rsp := v1.NewTmRegResponse(req.Id)
		rsp.Msg.(*pb.RegisterTMResponseProto).AbstractIdentifyResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
		logging.Debugf("send tm reg response %s", rsp.String())

		ctx := context.Background()

		return c.SendResponse(ctx, rsp)
	}
	return nil
}

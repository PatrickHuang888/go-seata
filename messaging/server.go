package messaging

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"go.uber.org/atomic"
	"net"
)

type Server struct {
	addr string

	ready chan struct{}

	close        atomic.Bool
	closing      chan struct{}
	closed       chan struct{}
	channelClose chan string

	channels map[string]*Channel

	syncReqHandlers  []ServerMessageHandler
	asyncReqHandlers []ServerMessageHandler
	asyncRspHandlers []ServerMessageHandler
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, close: atomic.Bool{}, closing: make(chan struct{}, 1), closed: make(chan struct{}, 1),
		channels: map[string]*Channel{}, channelClose: make(chan string), ready: make(chan struct{}, 1)}
}

func (s *Server) Serv() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		logging.Errorf("listen error ", err)
		s.Close()
	}
	logging.Infof("listen on %s\n", s.addr)

	s.ready <- struct{}{}

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
		conn, err := l.Accept()

		if err != nil {
			if s.close.Load() {
				logging.Infof("closing listener")
				break loop
			} else {
				logging.Errorf("accept error %s\n", err.Error())
				continue
			}

		}

		config := DefaultConfig()
		config.EnableHeartBeat = false
		ch := NewChannelWithConfig("server", conn, config)
		ch.closeListener = s
		s.channels[ch.name] = ch

		h := &defaultChannelHandler{ch, s}
		ch.RegisterSyncRequestHandler(h)
		ch.RegisterAsyncResponseHandler(h)
		ch.RegisterAsyncRequestHandler(h)

		go ch.run()
	}

	fmt.Println("server closed")
	<-s.closed
}

type defaultChannelHandler struct {
	c   *Channel
	svr *Server
}

func (dch *defaultChannelHandler) HandleMessage(msg v1.Message) error {
	switch msg.Tp {
	case v1.MsgTypeRequestSync:
		for _, h := range dch.svr.syncReqHandlers {
			if err := h.HandleMessage(dch.c, msg); err != nil {
				logging.Errorf("%+v", err)
			}
		}

	case v1.MsgTypeHeartbeatRequest:
		fallthrough
	case v1.MsgTypeRequestOneway:
		for _, h := range dch.svr.asyncReqHandlers {
			if err := h.HandleMessage(dch.c, msg); err != nil {
				logging.Errorf("+v", err)
			}
		}

	case v1.MsgTypeResponse:
		for _, h := range dch.svr.asyncRspHandlers {
			if err := h.HandleMessage(dch.c, msg); err != nil {
				logging.Errorf("+v", err)
			}
		}

	default:
		logging.Errorf("message type unknown")
	}

	return nil
}

func (s *Server) RegisterSyncRequestHandler(h ServerMessageHandler) {
	s.syncReqHandlers = append(s.syncReqHandlers, h)
}

func (s *Server) RegisterAsyncRequestHandler(h ServerMessageHandler) {
	s.asyncReqHandlers = append(s.asyncReqHandlers, h)
}

func (s *Server) RegisterAsyncRspHandler(h ServerMessageHandler) {
	s.asyncRspHandlers = append(s.asyncRspHandlers, h)
}

func (s *Server) Close() {
	s.closing <- struct{}{}
}

func (s *Server) ChannelClose(name string) {
	s.channelClose <- name
}

func (s *Server) Ready() <-chan struct{} {
	return s.ready
}

type ServerMessageHandler interface {
	HandleMessage(c *Channel, msg v1.Message) error
}

package main

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"strconv"
	"time"
)

func main() {
	var wait chan struct{}

	s := messaging.NewServer("localhost:7788")
	s.RegisterSyncRequestHandler(&testHandler{})
	s.RegisterAsyncRequestHandler(&pingHandler{})
	go s.Serv()

	<-wait
}

type pingHandler struct {
}

func (h *pingHandler) HandleMessage(c *messaging.Channel, msg v1.Message) error {
	switch msg.Tp {
	case v1.MsgTypeHeartbeatRequest:
		fmt.Printf("received ping message id %d\n", msg.Id)
		rsp := v1.NewHeartbeatResponse(msg.Id)
		if err := c.SendResponse(context.Background(), &rsp); err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

type testHandler struct {
}

func (h *testHandler) HandleMessage(c *messaging.Channel, msg v1.Message) error {
	req, ok := msg.Msg.(*pb.TestRequestProto)
	if ok {
		switch req.GetType() {
		case pb.TestMessageType_Timeout:
			fallthrough
		case pb.TestMessageType_Deadline:
			fallthrough
		case pb.TestMessageType_Cancel:
			sleep, err := strconv.Atoi(req.GetParam1())
			rsp := v1.NewTestResponse(1)
			if err == nil {
				rsp.Msg.(*pb.TestResponseProto).AbstractIdentifyResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
			} else {
				logging.Errorf("test message param error %s", err)
			}
			time.Sleep(time.Duration(sleep) * time.Second)
			if err := c.SendResponse(context.Background(), &rsp); err != nil {
				logging.Debug(err)
			}

		default:
			rsp := v1.NewTestResponse(msg.Id)
			rsp.Msg.(*pb.TestResponseProto).AbstractIdentifyResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
			if err := c.SendResponse(context.Background(), &rsp); err != nil {
				logging.Debug(err)
			}
		}
	}
	return nil
}

package messaging

import (
	"context"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"go.uber.org/atomic"
	"strconv"
	"testing"
	"time"
)

func TestBasicSendAndReceive(t *testing.T) {

	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTmReg)
	go s.Serv()

	<-s.ready

	c, err := NewClient("localhost:7788")
	if err != nil {
		t.Fatal(err)
	}

	req := v1.NewTmRegRequest("tm-test", "tx-group-test")
	rsp, err := c.Call(req)
	if err != nil {
		t.Fatal(err)
	}
	tmRsp, ok := rsp.Msg.(*pb.RegisterTMResponseProto)
	if !ok {
		t.Errorf("not tm register response")
	}

	if tmRsp.AbstractIdentifyResponse.AbstractResultMessage.GetResultCode() != pb.ResultCodeProto_Success {
		t.Errorf("result error")
	}

	c.Close()
	time.Sleep(2 * time.Second)
	s.Close()
	time.Sleep(2 * time.Second)
}

func handleTest(c *Channel, msg v1.Message) error {
	req, ok := msg.Msg.(*pb.TestRequestProto)
	if ok && (req.GetType() == pb.TestMessageType_Timeout || req.GetType() == pb.TestMessageType_Deadline ||
		req.GetType() == pb.TestMessageType_Cancel) {
		sleep, err := strconv.Atoi(req.GetParam1())
		rsp := newTestResponse()
		if err == nil {
			rsp.Msg.(*pb.TestResponseProto).AbstractIdentifyResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
		} else {
			logging.Errorf("test message param error %s", err)
		}
		time.Sleep(time.Duration(sleep) * time.Second)
		if err := c.SendResponse(context.Background(), &rsp); err != nil {
			logging.Debug(err)
		}
	}
	return nil
}

func TestTimeout(t *testing.T) {
	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTest)
	go s.Serv()

	<-s.ready

	c, err := NewClientWithConfig("localhost:7788", 500)
	if err != nil {
		t.Fatal(err)
	}

	req := newTestTimeoutRequest()

	_, err = c.Call(req)
	if err == nil {
		t.Fail()
	}

	c.Close()
	s.Close()
}

func TestDeadline(t *testing.T) {
	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTest)
	go s.Serv()

	<-s.ready

	c, err := NewClient("localhost:7788")
	if err != nil {
		t.Fatal(err)
	}

	req := newTestDeadlineRequest()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(500*time.Millisecond))

	var result atomic.Bool

	go func() {
		_, err = c.CallWithCtx(ctx, req)
		if err == nil {
			t.Fail()
		} else {
			logging.Debug(err)
		}
		result.CAS(false, true)
	}()

	time.Sleep(1 * time.Second)

	if !result.Load() {
		t.Fail()
	}

	c.Close()
	s.Close()
}

func TestCancel(t *testing.T) {
	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTest)
	go s.Serv()

	<-s.ready

	c, err := NewClient("localhost:7788")
	if err != nil {
		t.Fatal(err)
	}

	req := newTestDeadlineRequest()

	ctx, cancel := context.WithCancel(context.Background())

	var result atomic.Bool
	ready := make(chan struct{})
	wait := make(chan struct{})

	go func() {
		ready <- struct{}{}
		_, err = c.CallWithCtx(ctx, req)
		if err == nil {
			t.Fail()
		} else {
			logging.Debug(err)
		}
		result.CAS(false, true)
		wait <- struct{}{}
	}()

	<-ready
	cancel()
	<-wait

	if !result.Load() {
		t.Fail()
	}

	c.Close()
	s.Close()
}

func newTestDeadlineRequest() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeRequestSync, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestRequestProto{Type: pb.TestMessageType_Deadline, Param1: "50000"}}
}

func newTestTimeoutRequest() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeRequestSync, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestRequestProto{Type: pb.TestMessageType_Timeout, Param1: "50000"}}
}

func newTestResponse() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeResponse, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestResponseProto{AbstractIdentifyResponse: &pb.AbstractIdentifyResponseProto{
			AbstractResultMessage: &pb.AbstractResultMessageProto{}}}}
}

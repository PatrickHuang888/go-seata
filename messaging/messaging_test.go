package messaging

import (
	"context"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"go.uber.org/atomic"
	"strconv"
	"sync"
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

func TestTimeout(t *testing.T) {
	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTest)
	go s.Serv()

	<-s.ready

	config := DefaultConfig()
	config.Timeout = 500 * time.Millisecond
	c, err := NewClientWithConfig("localhost:7788", config)
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

func TestConcurrent(t *testing.T) {
	threads := 50
	results := make([]bool, threads)

	var waits sync.WaitGroup
	waits.Add(threads)

	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTest)
	go s.Serv()

	<-s.ready

	c, err := NewClient("localhost:7788")
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < threads; i++ {
		go func() {
			req := v1.NewTestRequest()
			rsp, err := c.Call(req)
			if err != nil {
				t.Fail()
			}
			if rsp.Msg.(*pb.TestResponseProto).AbstractIdentifyResponse.AbstractResultMessage.GetResultCode() !=
				pb.ResultCodeProto_Success {
				t.Fail()
			}
			results[rsp.Id-1] = true
			waits.Done()
		}()
	}

	waits.Wait()

	for _, r := range results {
		if !r {
			t.Fail()
		}
	}

	c.Close()
	s.Close()
}

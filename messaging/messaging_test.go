package messaging

import (
	"context"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
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

func handleTimeoutTest(c *Channel, msg v1.Message) error {
	req, ok := msg.Msg.(*pb.TestTimeoutRequestProto)
	if ok {
		sleep := req.GetSleepTime()
		time.Sleep(time.Duration(sleep) * time.Second)
		rsp := newTestTimeoutResponse()
		if err := c.SendResponse(context.Background(), &rsp); err != nil {
			logging.Debug(err)
		}
	}
	return nil
}

func TestTimeout(t *testing.T) {
	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTimeoutTest)
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

func newTestTimeoutRequest() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeRequestSync, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestTimeoutRequestProto{SleepTime: 50_000}}
}

func newTestTimeoutResponse() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeResponse, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestTimeoutResponseProto{}}
}

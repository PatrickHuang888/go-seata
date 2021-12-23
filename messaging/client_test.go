package messaging

import (
	"fmt"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"sync"
	"testing"
	"time"
)

// should start io.seata.core.rpc.netty.v1.ProtocolV1Server first
func TestCallToJava(t *testing.T) {
	c, err := NewClient("localhost:8811")
	if err != nil {
		t.Fatal(err)
	}

	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewTmRegisterRequest("tm-test", "tx-group-test")
	rsp, err := c.Call(req)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := rsp.Msg.(*pb.RegisterTMRequestProto)
	if !ok {
		t.Errorf("response not tm register request")
	}

	c.Close()
	//time.Sleep(5 * time.Second)
}

// should start io.seata.core.rpc.netty.v1.ProtocolV1Server first
func TestCallToJavaConcurrently(t *testing.T) {
	threads := 50
	var wg sync.WaitGroup
	wg.Add(threads)

	c, err := NewClient("localhost:8811")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			req := v1.NewSyncRequestMessage()
			req.Msg = v1.NewTmRegisterRequest("tm-test", "tx-group-test")
			rsp, err := c.Call(req)
			if err != nil {
				t.Fatal(err)
			}
			_, ok := rsp.Msg.(*pb.RegisterTMRequestProto)
			if !ok {
				t.Errorf("response not tm register request")
			}
		}()
	}

	wg.Wait()
	c.Close()
	//time.Sleep(5 * time.Second)
}

type testRmRspHandler struct {
	t    *testing.T
	wait chan struct{}
	c    *Client
}

func (h *testRmRspHandler) HandleMessage(msg v1.Message) error {
	_, ok := msg.Msg.(*pb.RegisterRMResponseProto)
	if !ok {
		h.t.Fatal("response type error")
	}
	fmt.Println("get the rm register response")
	h.wait <- struct{}{}
	return nil
}

// should connect to real seata-server
func TestAsyncCallToJava(t *testing.T) {
	wait := make(chan struct{})

	c, err := NewClient("localhost:8091")
	if err != nil {
		t.Fatal(err)
	}

	h := &testRmRspHandler{c: c, wait: wait, t: t}
	c.RegisterAsyncResponseHandler(h)

	req := v1.NewAsyncRequestMessage()
	req.Msg = v1.NewResourceRegisterRequest("rm-test", "tx-group-test", "resourceIds")
	if err = c.AsyncCall(req); err != nil {
		t.Fatalf("%+v", err)
	}

	<-wait
	fmt.Println("close client")
	c.Close()
	//time.Sleep(5 * time.Second)
}

type testPongHandler struct {
	result *bool
	c      *Client
}

func (h *testPongHandler) HandleMessage(msg v1.Message) error {
	if msg.Tp == v1.MsgTypeHeartbeatResponse {
		fmt.Println("get the heartbeat response")
		*h.result = true
	}
	return nil
}

// start seata server first
func TestPingToJava(t *testing.T) {
	var result bool

	config := DefaultConfig()
	config.WriteIdle = 1 * time.Second
	c, err := NewClientWithConfig("localhost:8091", "test-app", "tx-group", config)
	if err != nil {
		t.Fatal(err)
	}

	h := &testPongHandler{c: c, result: &result}
	c.RegisterAsyncResponseHandler(h)

	<-time.After(c.config.WriteIdle * 3) // read idle

	if !result {
		t.Fatal("no heartbeat response")
	}

	fmt.Println("close client")
	c.Close()
}

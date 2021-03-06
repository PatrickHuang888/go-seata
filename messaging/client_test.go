package messaging

import (
	"fmt"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"sync"
	"testing"
	"time"
)

var (
	javaServer = "localhost:8091"
	goServer   = "localhost:7788"
)

// should start io.seata.core.rpc.netty.v1.ProtocolV1Server first
func TestCallToJava(t *testing.T) {
	c, err := NewClient("localhost:8811", "test-client", "test-txGroup")
	if err != nil {
		t.Fatal(err)
	}

	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewTmRegisterRequest(c.appId, c.txGroup)
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

	c, err := NewClient("localhost:8811", "test-client", "test-txGroup")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			req := v1.NewSyncRequestMessage()
			req.Msg = v1.NewTmRegisterRequest(c.appId, c.txGroup)
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

// start seata-server first
func TestAsyncCallToJava(t *testing.T) {
	wait := make(chan struct{})

	c, err := NewClient(javaServer, "test-client", "test-txGroup")
	if err != nil {
		t.Fatal(err)
	}

	h := &testRmRspHandler{c: c, wait: wait, t: t}
	c.RegisterAsyncResponseHandler(h)

	req := v1.NewAsyncRequestMessage()
	req.Msg = v1.NewResourceRegisterRequest(c.appId, c.txGroup, "resourceIds")
	if err = c.AsyncCall(req); err != nil {
		t.Fatalf("%+v", err)
	}

	<-wait
	fmt.Println("close client")
	c.Close()
	//time.Sleep(5 * time.Second)
}

type testPongHandler struct {
	result int
	c      *Client
}

func (h *testPongHandler) HandleMessage(msg v1.Message) error {
	if msg.Tp == v1.MsgTypeHeartbeatResponse {
		fmt.Println("get the heartbeat response")
		h.result++
	}
	return nil
}

// start server first
func TestPing(t *testing.T) {
	config := DefaultConfig()
	config.WriteIdle = 2 * time.Second
	config.Timeout = 7 * time.Second
	c, err := NewClientWithConfig(javaServer, "test-app", "tx-group", config)
	if err != nil {
		t.Fatal(err)
	}

	h := &testPongHandler{c: c}
	c.RegisterAsyncResponseHandler(h)

	<-time.After(c.config.Timeout) // read idle

	if h.result != 3 {
		t.Fatal("heartbeat response error")
	}

	fmt.Println("close client")
	c.Close()
}

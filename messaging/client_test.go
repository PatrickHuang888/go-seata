package messaging

import (
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"sync"
	"testing"
)

// should start io.seata.core.rpc.netty.v1.ProtocolV1Server first
func TestCallToJava(t *testing.T) {
	c, err := NewClient("localhost:8811")
	if err != nil {
		t.Fatal(err)
	}

	req := v1.NewTmRegRequest("tm-test", "tx-group-test")
	rsp, err := c.SyncCall(req)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := rsp.(*pb.RegisterTMRequestProto)
	if !ok {
		t.Errorf("response not tm register request")
	}

	msg := v1.NewRmRegMessage("rm-test", "tx-group-test", "resourceIds")
	if err =c.AsyncCall(msg);err!=nil {
		t.Fatalf("%+v", err)
	}

	c.Close()
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

			req := v1.NewTmRegRequest("req-test", "tx-group-test")
			rsp, err := c.SyncCall(req)
			if err != nil {
				t.Fatal(err)
			}
			_, ok := rsp.(*pb.RegisterTMRequestProto)
			if !ok {
				t.Errorf("response not tm register request")
			}
		}()
	}

	wg.Wait()
	c.Close()
}

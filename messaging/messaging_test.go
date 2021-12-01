package messaging

import (
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"testing"
	"time"
)

func TestTimeOut (t *testing.T) {

	s := NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTmReg)
	go s.Serv()

	time.Sleep(2 *time.Second)

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
		t.Errorf("response not tm register request")
	}

	if tmRsp.AbstractIdentifyResponse.AbstractResultMessage.GetResultCode()!= pb.ResultCodeProto_Success {
		t.Errorf("result error")
	}

	c.Close()
	time.Sleep(5 *time.Second)
	s.Close()
	time.Sleep(5 *time.Second)
}

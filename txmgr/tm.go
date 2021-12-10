package txmgr

import (
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
)

type TM struct {
	c *messaging.Client
}

func NewTm(svrAddr string) (*TM, error) {
	c, err := messaging.NewClient(svrAddr)
	if err != nil {
		return nil, err
	}
	return &TM{c: c}, nil
}

func (tm *TM) Register() error {
	msg := v1.NewTmRegRequest("go-client", "go-client-txgroup")

	rsp, err := tm.c.Call(msg)
	if err != nil {
		return errors.WithStack(err)
	}

	tmRegRsp, ok := rsp.Msg.(*pb.RegisterTMResponseProto)
	if !ok {
		return errors.New("not tm reg response")
	}
	if tmRegRsp.AbstractIdentifyResponse.AbstractResultMessage.GetResultCode() != pb.ResultCodeProto_Success {
		return errors.New("tm register failed")
	}

	return nil
}

func (tm *TM) Close() {
	tm.c.Close()
}

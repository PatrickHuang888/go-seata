package tm

import (
	"context"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
)

type TM struct {
	c *messaging.Client
}

func NewTm(c *messaging.Client) *TM{
	return &TM{c:c}
}

func (tm *TM) Register() error {
	msg := v1.NewTmRegRequest("go-client", "go-client-txgroup")

	ctx := context.Background()

	rsp, err := tm.c.Call(ctx, msg)
	if err != nil {
		return errors.WithStack(err)
	}

	tmRegRsp, ok := rsp.(*pb.RegisterTMResponseProto)
	if !ok {
		return errors.New("not tm reg response")
	}
	if tmRegRsp.AbstractIdentifyResponse.AbstractResultMessage.GetResultCode() != pb.ResultCodeProto_Success {
		return errors.New("tm register failed")
	}

	return nil
}


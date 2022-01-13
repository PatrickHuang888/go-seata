package txmgr

import (
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
	"time"
)

const (
	DefaultGlobalTxTimeout = 60 * time.Second
)

type TM struct {
	c *messaging.Client
}

func NewTm(c *messaging.Client) *TM {
	return &TM{c: c}
}

func (tm *TM) Register() error {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewTmRegisterRequest(tm.c.AppId(), tm.c.TxGroup())
	rsp, err := tm.c.Call(req)
	if err != nil {
		return errors.WithStack(err)
	}

	tmRegRsp, ok := rsp.Msg.(*pb.RegisterTMResponseProto)
	if !ok {
		return errors.New("not tm reg response")
	}
	if tmRegRsp.AbstractIdentifyResponse.AbstractResultMessage.GetResultCode() != pb.ResultCodeProto_Success {
		return errors.New(tmRegRsp.AbstractIdentifyResponse.AbstractResultMessage.Msg)
	}

	// todo: on register success

	return nil
}

// Begin name from Annotation in Java
func (tm *TM) Begin(name string, timeout time.Duration) (xid string, err error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewGlobalBeginRequest(name, int32(timeout.Milliseconds()))
	rsp, err := tm.c.Call(req)
	if err != nil {
		return "", err
	}
	result, ok := rsp.Msg.(*pb.GlobalBeginResponseProto)
	if !ok {
		return "", errors.New("response type error")
	}
	if result.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return result.Xid, errors.New(result.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return result.Xid, nil
}

func (tm *TM) Commit(xid string) (pb.GlobalStatusProto, error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewGlobalCommitRequest(xid)
	rsp, err := tm.c.Call(req)
	if err != nil {
		return pb.GlobalStatusProto_UnKnown, err
	}
	commitRsp, ok := rsp.Msg.(*pb.GlobalCommitResponseProto)
	if !ok {
		return pb.GlobalStatusProto_UnKnown, errors.New("response type error")
	}
	if commitRsp.AbstractGlobalEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return commitRsp.AbstractGlobalEndResponse.GlobalStatus, errors.New(commitRsp.AbstractGlobalEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return commitRsp.AbstractGlobalEndResponse.GlobalStatus, nil
}

func (tm *TM) Rollback(xid string) (pb.GlobalStatusProto, error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewGlobalRollbackRequest(xid)
	rsp, err := tm.c.Call(req)
	if err != nil {
		return pb.GlobalStatusProto_UnKnown, err
	}
	rollbackRsp, ok := rsp.Msg.(*pb.GlobalRollbackResponseProto)
	if !ok {
		return pb.GlobalStatusProto_UnKnown, errors.New("response type error")
	}
	if rollbackRsp.AbstractGlobalEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return rollbackRsp.AbstractGlobalEndResponse.GlobalStatus, errors.New(rollbackRsp.AbstractGlobalEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return rollbackRsp.AbstractGlobalEndResponse.GlobalStatus, nil
}

func (tm *TM) Close() {
	tm.c.Close()
}

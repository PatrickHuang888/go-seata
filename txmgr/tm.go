package txmgr

import (
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
	"time"
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

func (tm *TM) Begin(name string, timeout time.Duration) (xid string, err error) {
	begin := v1.NewGlobalBeginRequest(name, int32(timeout.Milliseconds()))
	rsp, err := tm.c.Call(begin)
	if err != nil {
		return "", err
	}
	result, ok := rsp.Msg.(*pb.GlobalBeginResponseProto)
	if ok {
		if result.AbstractTransactionResponse.AbstractResultMessage.ResultCode == pb.ResultCodeProto_Success {
			return result.Xid, nil
		} else {
			return "", errors.New("begin failed")
		}
	} else {
		return "", errors.New("response type error")
	}
}

func (tm *TM) Commit(xid string) (pb.GlobalStatusProto, error) {
	commit := v1.NewGlobalCommitRequest(xid)
	rsp, err := tm.c.Call(commit)
	if err != nil {
		return pb.GlobalStatusProto_UnKnown, err
	}
	commitRsp, ok := rsp.Msg.(*pb.GlobalCommitResponseProto)
	if ok {
		return commitRsp.AbstractGlobalEndResponse.GlobalStatus, nil
	} else {
		return pb.GlobalStatusProto_UnKnown, errors.New("response type error")
	}
}

func (tm *TM) Rollback(xid string) (pb.GlobalStatusProto, error) {
	rollback := v1.NewGlobalRollbackRequest(xid)
	rsp, err := tm.c.Call(rollback)
	if err != nil {
		return pb.GlobalStatusProto_UnKnown, err
	}
	commitRsp, ok := rsp.Msg.(*pb.GlobalRollbackResponseProto)
	if ok {
		return commitRsp.AbstractGlobalEndResponse.GlobalStatus, nil
	} else {
		return pb.GlobalStatusProto_UnKnown, errors.New("response type error")
	}
}

func (tm *TM) Close() {
	tm.c.Close()
}

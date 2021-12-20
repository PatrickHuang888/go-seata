package resmgr

import (
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
)

type RM struct {
	c *messaging.Client
}

func NewRM(c *messaging.Client) *RM {
	c.RegisterAsyncRspHandler(resourceRegistered)
	return &RM{c}
}

func resourceRegistered(c *messaging.Channel, msg v1.Message) error {
	rsp, ok := msg.Msg.(*pb.RegisterRMResponseProto)
	if ok {
		if rsp.AbstractIdentifyResponse.AbstractResultMessage.ResultCode == pb.ResultCodeProto_Success {
			logging.Infof("resource register success %s", rsp.String())
		} else {
			logging.Errorf("resource register error on server %s", rsp.AbstractIdentifyResponse.AbstractResultMessage.Msg)
		}
	}
	return nil
}

func (rm *RM) RegisterResource(tp pb.BranchTypeProto, resourceId string) {
	// todo: client reconnect

	req := v1.NewAsyncRequestMessage()
	req.Msg = v1.NewResourceRegisterRequest(rm.c.AppId(), rm.c.TxGroup(), resourceId)
	if err := rm.c.AsyncCall(req); err != nil {
		logging.Errorf("resource register error %+v", err)
	}
}

func (rm *RM) RegisterBranch(tp pb.BranchTypeProto, xid string, resourceId string) (branchId int64, err error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewBranchRegisterRequest(tp, xid, resourceId)

	rsp, err := rm.c.Call(req)
	if err != nil {
		return 0, err
	}
	regRsp, ok := rsp.Msg.(*pb.BranchRegisterResponseProto)
	if !ok {
		return 0, errors.New("response type error")
	}
	if regRsp.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return regRsp.BranchId, errors.Errorf("branch register fail %s", regRsp.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return regRsp.BranchId, nil
}

/*func (rm *RM) BranchCommit(tp pb.BranchTypeProto, branchId int64, xid string, resourceId string) (pb.BranchStatusProto, error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewBranchCommitRequest(tp, branchId, xid, resourceId)

	rsp, err := rm.c.Call(req)
	if err != nil {
		return pb.BranchStatusProto_BUnknown, err
	}

	commitRsp, ok := rsp.Msg.(*pb.BranchCommitResponseProto)
	if !ok {
		return pb.BranchStatusProto_BUnknown, errors.New("response type error")
	}
	if commitRsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return commitRsp.AbstractBranchEndResponse.BranchStatus,
			errors.New(commitRsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return commitRsp.AbstractBranchEndResponse.BranchStatus, nil
}

func (rm *RM) BranchRollback(tp pb.BranchTypeProto, branchId int64, xid string, resourceId string) (pb.BranchStatusProto, error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewBranchRollbackRequest(tp, branchId, xid, resourceId)

	rsp, err := rm.c.Call(req)
	if err != nil {
		return pb.BranchStatusProto_BUnknown, err
	}

	rollback, ok := rsp.Msg.(*pb.BranchRollbackResponseProto)
	if !ok {
		return pb.BranchStatusProto_BUnknown, errors.New("response type error")
	}
	if rollback.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode != pb.ResultCodeProto_Success {
		return rollback.AbstractBranchEndResponse.BranchStatus, errors.Errorf("branch rollback fail %s",
			rollback.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg)
	}
	return rollback.AbstractBranchEndResponse.BranchStatus, nil
}
*/

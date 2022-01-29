package resmgr

import (
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
)

type RM struct {
	*messaging.Client
}

func NewRM(c *messaging.Client) *RM {
	rm := &RM{c}
	c.RegisterAsyncResponseHandler(&rsRegHandler{})
	return rm
}

type rsRegHandler struct {
}

func (h *rsRegHandler) HandleMessage(msg v1.Message) error {
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
	req.Msg = v1.NewResourceRegisterRequest(rm.AppId(), rm.TxGroup(), resourceId)
	if err := rm.AsyncCall(req); err != nil {
		logging.Errorf("resource register error %+v", err)
	}
}

func (rm *RM) RegisterBranch(tp pb.BranchTypeProto, xid string, resourceId string) (branchId int64, err error) {
	req := v1.NewSyncRequestMessage()
	req.Msg = v1.NewBranchRegisterRequest(tp, xid, resourceId)

	rsp, err := rm.Call(req)
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

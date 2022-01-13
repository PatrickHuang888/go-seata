package resmgr

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
)

const (
	TccKey         = "TccKey"
	TccKeyBranchId = "BranchId"
	TccKeyXid      = "XId"
)

type TCCResource struct {
	resource

	service TCC
}

type TccRM struct {
	reses map[string]TCCResource
	*RM
}

func NewTccRM(c *messaging.Client) *TccRM {
	rm := &TccRM{
		reses: make(map[string]TCCResource),
		RM:    NewRM(c),
	}
	h := &tccMsgHandler{rm}
	c.RegisterSyncRequestHandler(h)
	return rm
}

func (rm *TccRM) RegisterResource(res TCCResource) error {
	rm.reses[res.Id()] = res
	rm.RM.RegisterResource(pb.BranchTypeProto_TCC, res.Id())
	return nil
}

type tccMsgHandler struct {
	rm *TccRM
}

func (h *tccMsgHandler) HandleMessage(msg v1.Message) error {
	switch msg.Msg.(type) {
	case *pb.BranchCommitRequestProto:
		logging.Debugf("handle branch commit %d", msg.Id)

		commit := msg.Msg.(*pb.BranchCommitRequestProto)
		commitRsp := h.rm.handleBranchCommit(commit)

		rsp := v1.NewResponseMessage(msg.Id)
		rsp.Msg = commitRsp
		ctx := context.Background()

		return h.rm.SendResponse(ctx, &rsp)

	case *pb.BranchRollbackRequestProto:
		fmt.Println("branch rollback")
	}

	return nil
}

func (rm *TccRM) handleBranchCommit(commit *pb.BranchCommitRequestProto) *pb.BranchCommitResponseProto {
	branchId := commit.AbstractBranchEndRequest.BranchId
	xid := commit.AbstractBranchEndRequest.Xid
	resourceId := commit.AbstractBranchEndRequest.ResourceId
	// todo: application data

	rsp := v1.NewBranchCommitResponse(xid, branchId)

	res, ok := rm.reses[resourceId]
	if !ok {
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = fmt.Sprintf("resource %s not found", resourceId)
		return rsp
	}

	actionCtx := make(map[string]interface{})
	actionCtx[TccKeyXid] = xid
	actionCtx[TccKeyBranchId] = branchId
	// refactoring: ctx passing
	ctx := context.WithValue(context.Background(), TccKey, actionCtx)
	if err := res.service.Commit(ctx); err != nil {
		// ?
		rsp.AbstractBranchEndResponse.BranchStatus = pb.BranchStatusProto_PhaseTwo_CommitFailed_Retryable
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = err.Error()
		return rsp
	}

	rsp.AbstractBranchEndResponse.BranchStatus = pb.BranchStatusProto_PhaseTwo_Committed
	rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
	return rsp
}

func (rm *TccRM) handleBranchRollback(rollback *pb.BranchRollbackRequestProto) *pb.BranchRollbackResponseProto {
	branchId := rollback.AbstractBranchEndRequest.BranchId
	xid := rollback.AbstractBranchEndRequest.Xid
	resourceId := rollback.AbstractBranchEndRequest.ResourceId
	// todo: application data

	rsp := v1.NewBranchRollbackResponse(xid, branchId)

	res, ok := rm.reses[resourceId]
	if !ok {
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = fmt.Sprintf("resource %s not found", resourceId)
		return rsp
	}

	actionCtx := make(map[string]interface{})
	actionCtx[TccKeyXid] = xid
	actionCtx[TccKeyBranchId] = branchId
	// refactoring: ctx passing
	ctx := context.WithValue(context.Background(), TccKey, actionCtx)
	if err := res.service.Rollback(ctx); err != nil {
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = err.Error()
		return rsp
	}

	rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
	return rsp
}

type TCC interface {
	Action(context.Context) error
	Commit(context.Context) error
	Rollback(context.Context) error
}

func GetTCCContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, TccKey, make(map[string]interface{}))
}

package resmgr

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
)

type DBResources struct {
	txs map[string]*sql.Tx
	rm  *RM
}

type xaResource struct {
	resourceId string
	xid        string
}

func NewDBResources(rm *RM) *DBResources {
	dbrs := &DBResources{rm: rm, txs: make(map[string]*sql.Tx)}
	h := &dbBranchMsgHandler{dbrs}
	dbrs.rm.RegisterSyncRequestHandler(h)
	return dbrs
}

type dbBranchMsgHandler struct {
	dbrs *DBResources
}

func (h *dbBranchMsgHandler) HandleMessage(msg v1.Message) error {
	rspMsg := v1.NewResponseMessage(msg.Id)

	switch msg.Msg.(type) {
	case *pb.BranchCommitRequestProto:
		commit := msg.Msg.(*pb.BranchCommitRequestProto)
		rsp := h.dbrs.handleBranchCommit(commit)
		rspMsg.Msg = rsp
		return h.dbrs.rm.SendResponse(context.Background(), &rspMsg)

	case *pb.BranchRollbackRequestProto:
		rollback := msg.Msg.(*pb.BranchRollbackRequestProto)
		rsp := h.dbrs.handleBranchRollback(rollback)
		rspMsg.Msg = rsp
		return h.dbrs.rm.SendResponse(context.Background(), &rspMsg)
	}

	return nil
}

func (dbrs *DBResources) handleBranchCommit(commit *pb.BranchCommitRequestProto) *pb.BranchCommitResponseProto {
	if commit.AbstractBranchEndRequest.BranchType != pb.BranchTypeProto_XA {
		logging.Errorf("branch type only support XA right now")
	}

	branchId := commit.AbstractBranchEndRequest.BranchId
	xid := commit.AbstractBranchEndRequest.Xid
	xaid := GetXaId(xid, branchId)

	rsp := v1.NewBranchCommitResponse(xid, branchId)
	rsp.AbstractBranchEndResponse.Xid = xid
	rsp.AbstractBranchEndResponse.BranchId = branchId

	tx, ok := dbrs.txs[xaid]
	if !ok {
		str := fmt.Sprintf("do not find tx %s", xaid)
		logging.Error(str)
		rsp.AbstractBranchEndResponse.BranchStatus = pb.BranchStatusProto_PhaseTwo_CommitFailed_Retryable
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = str
		return rsp
	}

	if _, err := tx.Exec(fmt.Sprintf("XA PREPARE '%s'", xaid)); err != nil {
		str := fmt.Sprintf("branch commit %s err", xaid)
		logging.Error(str)
		rsp.AbstractBranchEndResponse.BranchStatus = pb.BranchStatusProto_PhaseTwo_CommitFailed_Unretryable
		rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.Msg = err.Error()
		return rsp
	}

	rsp.AbstractBranchEndResponse.BranchStatus = pb.BranchStatusProto_PhaseTwo_Committed
	rsp.AbstractBranchEndResponse.AbstractTransactionResponse.AbstractResultMessage.ResultCode = pb.ResultCodeProto_Success
	return rsp
}

func (dbrs *DBResources) handleBranchRollback(rollback *pb.BranchRollbackRequestProto) *pb.BranchRollbackResponseProto {
	return nil
}

func GetXaId(xid string, branchId int64) string {
	return ""
}

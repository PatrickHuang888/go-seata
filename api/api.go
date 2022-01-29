package api

import (
	"context"
	"github.com/PatrickHuang888/go-seata/conf"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/PatrickHuang888/go-seata/resmgr"
)

const (
	TccKey         = "TccKey"
	TccKeyBranchId = "BranchId"
	TccKeyXid      = "XId"
	SeataXidKey    = "TX_XID"
)

type Resource interface {
	Id() string
	GroupId() string
	BranchType() pb.BranchTypeProto
}

type TCC interface {
	Action(ctx context.Context, req interface{}) (rsp interface{}, err error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type TCCResource interface {
	Resource
	TCC
}

func GetTCCResource(tcc TCC) TCCResource {
	return resmgr.NewTCCResource(conf.Config, tcc)
}

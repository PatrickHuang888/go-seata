package resmgr

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/api"
	"github.com/PatrickHuang888/go-seata/conf"
	"github.com/PatrickHuang888/go-seata/messaging"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/PatrickHuang888/go-seata/txmgr"
	"sync"
	"testing"
)

var serverAddr = "localhost:8091"
var appId = "test-client"
var txGroup = "test-txGroup"

func TestTCC(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(1)

	c, err := messaging.NewClient(serverAddr, appId, txGroup)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	tm := txmgr.NewTm(c)
	if err := tm.Register(); err != nil {
		t.Fatalf("%+v", err)
	}

	rm := NewTccRM(c)

	service := &testService{}
	tccResource := NewTCCResource(conf.Config, service)

	if err := rm.RegisterResource(tccResource); err != nil {
		t.Fatalf("%+v", err)
	}

	xid, err := tm.Begin("test-tx", txmgr.DefaultGlobalTxTimeout)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	branchId, err := rm.RegisterBranch(pb.BranchTypeProto_TCC, xid, tccResource.Id())
	if err != nil {
		t.Fatalf("%+v", err)
	}

	actionCtx := make(map[string]interface{})
	actionCtx[api.TccKeyXid] = xid
	actionCtx[api.TccKeyBranchId] = branchId
	ctx := context.WithValue(context.Background(), api.TccKey, actionCtx)
	_ = service.Action(ctx, nil)

	status, err := tm.Commit(xid)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if status != pb.GlobalStatusProto_Committed {
		t.Fail()
	}

	c.Close()
}

type testService struct {
}

func (s *testService) Action(ctx context.Context, req interface{}) error {
	fmt.Println("action")
	return nil
}

func (s *testService) Commit(ctx context.Context) error {
	params := ctx.Value(api.TccKey).(map[string]interface{})
	xid := params[api.TccKeyXid]

	fmt.Printf("commit xid %s\n", xid)
	return nil
}

func (s *testService) Rollback(ctx context.Context) error {
	fmt.Println("rollback")
	return nil
}

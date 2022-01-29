package interceptors

import (
	. "github.com/PatrickHuang888/go-seata/conf"
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	"github.com/PatrickHuang888/go-seata/resmgr"
	"github.com/PatrickHuang888/go-seata/txmgr"
)

var (
	tm     *txmgr.TM
	rm     *resmgr.RM
	tccRm  *resmgr.TccRM
	sagaRm *resmgr.SagaRM
)

func init() {
	c, err := messaging.NewClient(Config.SeataAddr, Config.AppId, Config.TxGroup)
	if err != nil {
		logging.Errorf("setup client to seata error %s", err.Error())
	}

	tm = txmgr.NewTm(c)
	tccRm = resmgr.NewTccRM(c)
}

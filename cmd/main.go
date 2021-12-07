package main

import (
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"time"
)

func main() {
	/*c, err := messaging.NewClient("localhost:8091")
	if err!=nil {
		fmt.Printf("%+v", err)
	}

	txmgr := tm.NewTm(c)
	err =txmgr.Register()
	if err!=nil {
		fmt.Printf(" tm reg error %+v", err)
	}*/

	var wait chan struct{}

	s := messaging.NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTimeoutTest)
	go s.Serv()

	<-s.Ready()

	<-wait
}

func handleTimeoutTest(c *messaging.Channel, msg v1.Message) error {
	req, ok := msg.Msg.(*pb.TestTimeoutRequestProto)
	if ok {
		sleep := req.GetSleepTime()
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	return nil
}

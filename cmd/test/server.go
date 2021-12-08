package main

import (
	"context"
	"fmt"
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/messaging"
	v1 "github.com/PatrickHuang888/go-seata/messaging/v1"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"time"
)

func main() {
	var wait chan struct{}

	s := messaging.NewServer("localhost:7788")
	s.RegisterRequestHandler(handleTimeoutTest)
	go s.Serv()

	<-wait
}

func handleTimeoutTest(c *messaging.Channel, msg v1.Message) error {
	req, ok := msg.Msg.(*pb.TestTimeoutRequestProto)
	if ok {
		sleep := time.Duration(req.GetSleepTime())
		s := sleep * time.Millisecond
		fmt.Printf("sleep %s\n", s.String())
		time.Sleep(s)
		rsp := newTestTimeoutResponse()
		if err := c.SendResponse(context.Background(), &rsp); err != nil {
			logging.Debug(err)
		}
	}
	return nil
}

func newTestTimeoutResponse() v1.Message {
	return v1.Message{Id: 1, Tp: v1.MsgTypeResponse, Ser: v1.SerializerProtoBuf, Ver: v1.Version,
		Msg: &pb.TestTimeoutResponseProto{}}
}

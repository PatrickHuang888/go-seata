package main

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/messaging"
)

func main() {
	client, err := messaging.NewClient("localhost:8091")
	if err!=nil {
		fmt.Printf("%+v", err)
	}
	msg := messaging.NewTmRegRequest("go-client", "tx-group")

	_, err=client.SendTmRegMsg(msg)
	if err!=nil {
		fmt.Printf("%+v", err)
	}

}

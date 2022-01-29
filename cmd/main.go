package main

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/messaging"
	"github.com/PatrickHuang888/go-seata/txmgr"
)

func main() {

	c, err := messaging.NewClient("localhost:8091", "appid", "txGroup")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	tm := txmgr.NewTm(c)
	err = tm.Register()
	if err != nil {
		fmt.Printf(" tm reg error %+v", err)
	}

	tm.Close()
}

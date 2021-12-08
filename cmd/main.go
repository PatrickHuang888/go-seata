package main

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/messaging"
	"github.com/PatrickHuang888/go-seata/tm"
)

func main() {
	c, err := messaging.NewClient("localhost:8091")
	if err != nil {
		fmt.Printf("%+v", err)
	}

	txmgr := tm.NewTm(c)
	err = txmgr.Register()
	if err != nil {
		fmt.Printf(" tm reg error %+v", err)
	}

}

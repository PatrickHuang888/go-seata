package main

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/txmgr"
)

func main() {

	tm, err := txmgr.NewTm("localhost:8091")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	err = tm.Register()
	if err != nil {
		fmt.Printf(" tm reg error %+v", err)
	}

	tm.Close()
}

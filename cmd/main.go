package main

import (
	"fmt"
	"github.com/PatrickHuang888/go-seata/messaging"
)

func main() {
	c, err := messaging.NewClient("localhost:8091")
	if err!=nil {
		fmt.Printf("%+v", err)
	}
	err =c.TmReg()
	if err!=nil {
		fmt.Printf(" tm reg error %+v", err)
	}
}
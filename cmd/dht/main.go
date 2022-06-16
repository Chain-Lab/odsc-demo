package main

import (
	"github.com/decision2016/osc/internal/node"
	"os"
	"os/signal"
)

func main() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)

	go node.GetDHTInstance()

	<-cs

	return
}

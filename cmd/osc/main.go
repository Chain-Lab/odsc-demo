package main

import (
	"github.com/decision2016/osc/internal/ethereum"
	"github.com/decision2016/osc/internal/rpc"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := cli.NewApp()

	app.Name = "osc"

	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "osc start --config <config path>",
			Action: startNode,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func startNode(c *cli.Context) (err error) {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)

	go rpc.StartRPCService()
	go ethereum.ListenEthereumContract()

	go func() {
		<-cs
		os.Exit(1)
	}()

	return
}

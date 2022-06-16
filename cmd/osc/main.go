package main

import (
	"github.com/decision2016/osc/internal/ethereum"
	"github.com/decision2016/osc/internal/node"
	"github.com/decision2016/osc/internal/rpc"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
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

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&utils.Formatter{})
	logrus.SetLevel(logrus.TraceLevel)

	go rpc.StartRPCService()
	go ethereum.ListenEthereumContract()
	go node.GetDHTInstance()

	if err != nil {
		return err
	}

	<-cs

	return
}

package rpc

import (
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartRPCService() {
	config := utils.ConfigInstance()
	port, err := config.Section("rpc").Key("port").Int()

	grpcServer := grpc.NewServer()
	RegisterCommandServiceServer(grpcServer, new(CommandService))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("Start RPC service on port ", port)

	grpcServer.Serve(listen)
}

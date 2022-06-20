package rpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func CallLocalRPC(command, subCommand string, params *map[string]*structpb.Value) {
	conn, err := grpc.Dial("", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	client := NewCommandServiceClient(conn)

	req := &CommandRequest{
		Command:    &command,
		SubCommand: &subCommand,
		Params:     *params,
	}

	reply, err := client.CommandProxy(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	logrus.Info(reply.GetStatus(), " ", reply.GetMsg())
}
package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func UploadKeystoreJson(keystore string) {
	conn, err := grpc.Dial("", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	client := NewCommandServiceClient(conn)

	command := "wallet"
	subCommand := "upload"

	req := &CommandRequest{
		Command:    &command,
		SubCommand: &subCommand,
		Params: Params{
			"keystore": &structpb.Value{Kind: &structpb.Value_StringValue{
				StringValue: keystore,
			}},
		},
	}

	reply, err := client.CommandProxy(context.Background(), req)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetStatus())

	return nil
}

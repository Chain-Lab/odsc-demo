package rpc

import (
	"context"
	"log"
)

type CommandService struct {
}

func (cs *CommandService) CommandProxy(ctx context.Context, req *CommandRequest) (resp *CommandRespond, err error) {
	command := req.GetCommand()
	subCommand := req.GetSubCommand()
	params := req.GetParams()

	log.Println(params["address"].GetStringValue())

	switch command {
	case "permission":
		resp = PermissionProxy(subCommand, &params)
	case "data":
		DataProxy(subCommand, &params, resp)
	case "genesis":
		GenesisProxy(resp)
	}

	return
}

func (cs *CommandService) mustEmbedUnimplementedCommandServiceServer() {
	//TODO implement me
	return
}

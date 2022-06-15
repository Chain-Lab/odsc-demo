package rpc

import (
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func PermissionProxy(subCommand string, params *map[string]*structpb.Value) (resp *CommandRespond) {
	switch subCommand {
	case "add":
		resp = subPermissionAdd(params)
	case "revoke":
		resp = subPermissionRevoke(params)
	}

	return
}

func subPermissionAdd(params *map[string]*structpb.Value) (resp *CommandRespond) {
	address := (*params)["address"].GetStringValue()
	log.Println("Receive command permission add --address ", address)

	resp = respWrite(0, "Success")
	return
}

func subPermissionRevoke(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

package rpc

import "google.golang.org/protobuf/types/known/structpb"

func PermissionProxy(subCommand string, params *map[string]*structpb.Value, resp *CommandRespond) {
	switch subCommand {
	case "add":
		subPermissionAdd(params, resp)
	case "revoke":
		subPermissionRevoke(params, resp)
	}
}

func subPermissionAdd(params *map[string]*structpb.Value, resp *CommandRespond) {

}

func subPermissionRevoke(params *map[string]*structpb.Value, resp *CommandRespond) {

}

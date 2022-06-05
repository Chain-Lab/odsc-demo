package rpc

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
)

type PermissionCommandService struct {
}

func (pcs *PermissionCommandService) PermissionProxy(req PermissionCommandRequest,
	resp *PermissionCommandRespond) (err error) {

	subCommand := req.GetSubCommand()
	params := req.GetParams()

	switch subCommand {
	// 传入指针避免构造
	case "add":
		subPermissionAdd(&params)
	case "revoke":
		subPermissionRevoke(&params)
	}

	return nil
}

func subPermissionAdd(params *map[string]*structpb.Value) {

}

func subPermissionRevoke(params *map[string]*structpb.Value) {

}

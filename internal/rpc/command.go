package rpc

import "google.golang.org/protobuf/types/known/structpb"

type Params map[string]*structpb.Value

type CommandService struct {
}

func (cs *CommandService) CommandProxy(req CommandRequest, resp *CommandRespond) error {

	command := req.GetCommand()
	subCommand := req.GetSubCommand()
	params := req.GetParams()

	switch command {
	case "permission":
		PermissionProxy(subCommand, &params, resp)
	case "certificate":
		CertificateProxy(subCommand, &params, resp)
	}

	return nil
}

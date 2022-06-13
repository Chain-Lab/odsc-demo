package rpc

type CommandService struct {
}

func (cs *CommandService) CommandProxy(req CommandRequest, resp *CommandRespond) error {

	command := req.GetCommand()
	subCommand := req.GetSubCommand()
	params := req.GetParams()

	switch command {
	case "permission":
		PermissionProxy(subCommand, &params, resp)
	case "data":
		DataProxy(subCommand, &params, resp)
	}

	return nil
}

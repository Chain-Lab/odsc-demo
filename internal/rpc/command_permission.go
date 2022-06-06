package rpc

func PermissionProxy(subCommand string, params *Params, resp *CommandRespond) {
	switch subCommand {
	case "add":
		subPermissionAdd(params, resp)
	case "revoke":
		subPermissionRevoke(params, resp)
	}
}

func subPermissionAdd(params *Params, resp *CommandRespond) {

}

func subPermissionRevoke(params *Params, resp *CommandRespond) {

}

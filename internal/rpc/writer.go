package rpc

func respWrite(status int32, msg string, resp *CommandRespond) {
	resp = &CommandRespond{
		Status: &status,
		Msg:    &msg,
	}
}

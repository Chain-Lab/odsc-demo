package rpc

import "google.golang.org/protobuf/types/known/structpb"

func DataProxy(subCommand string, params *map[string]*structpb.Value, resp *CommandRespond) {
	switch subCommand {
	case "publish":
		subDataPublish(params, resp)
	case "verify":
		subDataVerify(params, resp)
	case "modify":
		subDataModify(params, resp)
	case "invoke":
		subDataInvoke(params, resp)
	}
}

// subDataPublish 命令 "app data publish" 的执行入口
func subDataPublish(params *map[string]*structpb.Value, resp *CommandRespond) {

}

// subDataVerify 命令 "app data verify" 的执行入口
func subDataVerify(params *map[string]*structpb.Value, resp *CommandRespond) {

}

// subDataModify 命令 "app data modify" 的执行入口
func subDataModify(params *map[string]*structpb.Value, resp *CommandRespond) {

}

// subDataInvoke 命令 "app data invoke" 的执行入口
func subDataInvoke(params *map[string]*structpb.Value, resp *CommandRespond) {

}

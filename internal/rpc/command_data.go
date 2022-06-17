package rpc

import (
	"github.com/decision2016/osc/internal/node"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

func DataProxy(subCommand string, params *map[string]*structpb.Value) (resp *CommandRespond) {
	switch subCommand {
	case "publish":
		resp = subDataPublish(params)
	case "verify":
		resp = subDataVerify(params)
	case "modify":
		resp = subDataModify(params)
	case "invoke":
		resp = subDataInvoke(params)
	}
	return
}

// subDataPublish 命令 "app data publish" 的执行入口
func subDataPublish(params *map[string]*structpb.Value) (resp *CommandRespond) {
	key := (*params)["key"].GetStringValue()
	value := (*params)["value"].GetStringValue()

	logrus.Info("Receive request key: ", key, " value: ", value)

	err := node.PutDataToNetwork(key, []byte(value))

	if err != nil {
		logrus.Warn(err)
		resp = respWrite(-1, "")
		return
	}

	resp = respWrite(0, "")
	return
}

// subDataVerify 命令 "app data verify" 的执行入口
func subDataVerify(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

// subDataModify 命令 "app data modify" 的执行入口
func subDataModify(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

// subDataInvoke 命令 "app data invoke" 的执行入口
func subDataInvoke(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

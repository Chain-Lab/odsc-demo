package node

import "github.com/decision2016/osc/internal/utils"

var config = utils.ConfigInstance()

func Start() (err error) {
	runAsBootstrap, err := config.Section("node").Key("run_as_bootstrap").Bool()


	if err != nil {
		return
	}
}
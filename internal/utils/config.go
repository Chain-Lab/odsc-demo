package utils

import (
	"gopkg.in/ini.v1"
	"sync"
)

var cfgInstance *ini.File
var configOnce sync.Once

func ConfigInstance() *ini.File {
	configOnce.Do(func() {
		cfgInstance, _ = ini.Load("conf.ini")
	})
	return cfgInstance
}

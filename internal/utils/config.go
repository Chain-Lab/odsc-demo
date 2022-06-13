package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"sync"
)

var cfgInstance *ini.File
var configOnce sync.Once

func ConfigInstance() *ini.File {
	var err error
	configOnce.Do(func() {
		cfgInstance, err = ini.Load("conf.ini")
		if err != nil {
			log.Fatal(err)
		}
	})
	return cfgInstance
}

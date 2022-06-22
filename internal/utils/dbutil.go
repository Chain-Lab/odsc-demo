package utils

import (
	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
	"sync"
)

var dbInstance *kivik.Client = nil
var dbOnce sync.Once

func DBInstance() (client *kivik.Client, err error) {
	dbOnce.Do(func() {
		config := ConfigInstance()
		dbUrl := config.Section("couch").Key("address").String()
		//dbUrl := "http://admin:long2018nian!@42.193.15.205:5984/"

		dbInstance, err = kivik.New("couch", dbUrl)

		if err != nil {
			return
		}
	})

	return dbInstance, nil
}

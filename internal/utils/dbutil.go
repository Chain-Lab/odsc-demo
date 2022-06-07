package utils

import (
	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
)

var dbInstance *kivik.Client = nil

func DBInstance() (client *kivik.Client, err error) {
	if dbInstance == nil {
		config := ConfigInstance()
		dbUrl := config.Section("couch").Key("address").String()

		dbInstance, err = kivik.New("couch", dbUrl)

		if err != nil {
			return
		}
	}

	return dbInstance, nil
}

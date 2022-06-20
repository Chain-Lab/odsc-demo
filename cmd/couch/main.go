package main

import (
	"context"
	"github.com/decision2016/osc/internal/utils"
)

func main() {
	dbutil, _ := utils.DBInstance()

	db := dbutil.DB("data")
	data := db.Get(context.TODO(), "436e993d72e1f1e509afbbb44b5ef00885f1906087621b114a7a7b2f39da035d")

	var genesisData utils.GenesisData

	data.ScanDoc(&genesisData)

	println(genesisData.Random)
}

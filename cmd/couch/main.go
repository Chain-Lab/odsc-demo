package main

import (
	"context"
	"github.com/decision2016/osc/internal/utils"
)

func main() {
	dbutil, _ := utils.DBInstance()

	db, _ := dbutil.DBExists(context.Background(), "keystore")

	println(db)
}

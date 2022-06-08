package rpc

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/decision2016/osc/internal/utils"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func WalletProxy(subCommand string, params *map[string]*structpb.Value, resp *CommandRespond) {
	switch subCommand {
	// 上传已加密的钱包
	case "upload":
		subWalletUpload(params, resp)
	}
}

func subWalletUpload(params *map[string]*structpb.Value, resp *CommandRespond) {
	keystoreJson := (*params)["keystore"].String()

	identityKeyBytes := sha256.Sum256([]byte(keystoreJson))
	identityKey := hex.EncodeToString(identityKeyBytes[:])

	dbutil, err := utils.DBInstance()

	if err != nil {
		log.Fatal(err)
		return
	}

	dbExists, err := dbutil.DBExists(context.Background(), "keystore")

	if err != nil {
		log.Fatal(err)
		return
	}

	if !dbExists {
		err := dbutil.CreateDB(context.Background(), "keysotre")

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	db := dbutil.DB("keystore")

	doc := map[string]interface{} {
		"_id": identityKey,
		"keystore": keystoreJson,
	}

	rev, err := db.Put(context.TODO(), identityKey, doc)

	if err != nil {
		log.Fatal(err)

		resp = &CommandRespond{}
		resp.
	}
}

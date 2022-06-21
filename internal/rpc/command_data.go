package rpc

import (
	"context"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/decision2016/go-crypto/ecdcs"
	"github.com/decision2016/osc/internal/ethereum"
	"github.com/decision2016/osc/internal/node"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

func DataProxy(subCommand string, params *map[string]*structpb.Value) (resp *CommandRespond) {
	switch subCommand {
	case "publish":
		resp = subDataPublish(params)
	case "verify":
		resp = subDataVerify(params)
	case "modify":
		resp = subDataModify(params)
	case "revoke":
		resp = subDataRevoke(params)
	}
	return
}

// subDataPublish 命令 "app data publish" 的执行入口
func subDataPublish(params *map[string]*structpb.Value) (resp *CommandRespond) {
	data := (*params)["data"].GetStringValue()

	bytesData, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		resp = respWrite(-1, "Base64 decode failed.")
	}

	config := utils.ConfigInstance()
	localAddress := config.Section("wallet").Key("address").String()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	chameleonSec := ecdcs.PrivateKey{}
	err = chameleonSec.FromHexString(hexPrivateKey)
	chameleonPub, err := chameleonSec.ExportPublicKey(elliptic.P256())

	if err != nil {
		resp = respWrite(-1, "Parse private key error.")
		return
	}

	blockHash, err := ethereum.GetLatestBlockHash()
	if err != nil {
		resp = respWrite(-1, "Get latest block hash error.")
		return
	}

	newData := utils.StorageData{
		Id:        "",
		Random:    "",
		CreatedTx: "",
		Version:   0,
		Operator:  localAddress,
		Timestamp: time.Now().Unix(),
		Data:      bytesData,
		BlockHash: blockHash,
		PublicKey: chameleonPub.ToHexString(),
	}

	newDataJson, err := json.Marshal(newData)

	if err != nil {
		resp = respWrite(-1, "Parse data to json failed.")
		return
	}

	random, signature, err := chameleonSec.Signature(elliptic.P256(), newDataJson)

	if err != nil {
		resp = respWrite(-1, "Chameleon signature failed.")
		return
	}

	newData.Random = hex.EncodeToString(random)
	newData.Id = hex.EncodeToString(signature)

	tx, err := ethereum.CallContract(ethereum.CREATE, signature, hex.EncodeToString(random))

	if err != nil {
		resp = respWrite(-1, "Call ethereum contract failed")
		logrus.Error(err)
		return
	}

	newData.CreatedTx = tx
	newDataJson, err = json.Marshal(newData)

	if err != nil {
		resp = respWrite(-1, "Parse data to json failed.")
		return
	}

	dbUtil, err := utils.DBInstance()
	if err != nil {
		respWrite(-1, "Get database instance failed.")
		logrus.Error(err)
		return
	}

	db := dbUtil.DB("data")
	rev, err := db.Put(context.Background(), newData.Id, newDataJson)
	if err != nil {
		resp = respWrite(-1, "Write database error.")
		logrus.Error(err)
		return
	}

	err = node.PutDataToNetwork(newData.Id, newDataJson)

	if err != nil {
		resp = respWrite(-1, "Put data to P2P network failed.")
		logrus.Error(err)
		return
	}

	resp = respWrite(0, rev)
	return
}

// subDataVerify 命令 "app data verify" 的执行入口
func subDataVerify(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

// subDataModify 命令 "app data modify" 的执行入口
func subDataModify(params *map[string]*structpb.Value) (resp *CommandRespond) {
	data := (*params)["data"].GetStringValue()
	key := (*params)["key"].GetStringValue()

	bytesData, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		resp = respWrite(-1, "Base64 decode failed.")
	}

	dbUtil, err := utils.DBInstance()
	if err != nil {
		resp = respWrite(-1, "Get database instance failed.")
		return
	}

	db := dbUtil.DB("data")
	row := db.Get(context.TODO(), key)

	config := utils.ConfigInstance()
	//localAddress := config.Section("wallet").Key("address").String()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	chameleonSec := ecdcs.PrivateKey{}
	err = chameleonSec.FromHexString(hexPrivateKey)
	//chameleonPub, err := chameleonSec.ExportPublicKey(elliptic.P256())

	if err != nil {
		resp = respWrite(-1, "Parse private key error.")
		return
	}

	var storageData utils.StorageData
	err = row.ScanDoc(&storageData)
	if err != nil {
		resp = respWrite(-1, "Fetch data error.")
		logrus.Error(err)
		return
	}
	chameleonSignature := storageData.Id
	originRandom, err := hex.DecodeString(storageData.Random)

	if err != nil {
		resp = respWrite(-1, "Decode chameleon  data error.")
		logrus.Error(err)
		return
	}

	storageData.Id = ""
	storageData.CreatedTx = ""
	storageData.Random = ""

	originData, err := json.Marshal(storageData)

	storageData.Version += 1
	storageData.Data = bytesData

	dataJson, err := json.Marshal(storageData)

	if err != nil {
		resp = respWrite(-1, "Data serialize failed.")
		logrus.Error(err)
		return
	}

	newRandom := chameleonSec.ReSignature(elliptic.P256(), originData, dataJson, originRandom)

	storageData.Id = chameleonSignature
	storageData.Random = hex.EncodeToString(newRandom)
	hexChameleonSignature, err := hex.DecodeString(chameleonSignature)

	if err != nil {
		resp = respWrite(-1, "Decode chameleon signature failed.")
		logrus.Error(err)
		return
	}

	tx, err := ethereum.CallContract(ethereum.MODIFY, hexChameleonSignature, hex.EncodeToString(newRandom))

	if err != nil {
		resp = respWrite(-1, "Call ethereum contract failed")
		logrus.Error(err)
		return
	}

	storageData.CreatedTx = tx
	newRev, err := db.Put(context.TODO(), chameleonSignature, storageData)
	storageDataJson, err := json.Marshal(storageData)

	if err != nil {
		resp = respWrite(-1, "Storage data serialize failed.")
		logrus.Error(resp)
		return
	}

	err = node.PutDataToNetwork(storageData.Id, storageDataJson)

	if err != nil {
		resp = respWrite(-1, "Put data to P2P network failed.")
		logrus.Error(err)
		return
	}

	resp = respWrite(0, newRev)
	return
}

// subDataRevoke 命令 "app data invoke" 的执行入口
func subDataRevoke(params *map[string]*structpb.Value) (resp *CommandRespond) {
	return
}

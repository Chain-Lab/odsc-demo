package rpc

import (
	"context"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"github.com/decision2016/go-crypto/ecdcs"
	"github.com/decision2016/osc/internal/ethereum"
	"github.com/decision2016/osc/internal/node"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

func PermissionProxy(subCommand string, params *map[string]*structpb.Value) (resp *CommandRespond) {
	switch subCommand {
	case "add":
		resp = subPermissionAdd(params)
	case "revoke":
		resp = subPermissionRevoke(params)
	}

	return
}

func subPermissionAdd(params *map[string]*structpb.Value) (resp *CommandRespond) {
	address := (*params)["address"].GetStringValue()

	genesisInt64, err := ethereum.ReadGenesis()
	genesisIndex := hex.EncodeToString(genesisInt64.Bytes())
	if err != nil {
		resp = respWrite(-1, "Read genesis index from contract failed.")
		logrus.Error(err)
		return
	}

	config := utils.ConfigInstance()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	chameleonSec := ecdcs.PrivateKey{}
	err = chameleonSec.FromHexString(hexPrivateKey)

	if err != nil {
		resp = respWrite(-1, "Read config file failed.")
		logrus.Error(err)
		return
	}

	dbUtil, err := utils.DBInstance()

	if err != nil {
		resp = respWrite(-1, "Get database instance failed.")
		logrus.Error(err)
		return
	}

	db := dbUtil.DB("data")

	var genesisData utils.GenesisData
	row := db.Get(context.TODO(), genesisIndex)
	err = row.ScanDoc(&genesisData)

	list := genesisData.PermissionList

	if addressInList(&list, address) {
		resp = respWrite(-1, "Address already in list.")
		return
	}

	list = append(list, address)
	hexKey := genesisData.Id
	hexRandom := genesisData.Random

	originRandom, err := hex.DecodeString(hexRandom)

	if err != nil {
		resp = respWrite(-1, "Decode chameleon  data error.")
		logrus.Error(err)
		return
	}

	originKey, err := hex.DecodeString(hexKey)
	if err != nil {
		resp = respWrite(-1, "Decode chameleon  data error.")
		logrus.Error(err)
		return
	}

	genesisData.Id = ""
	genesisData.Random = ""
	genesisData.CreatedTx = ""

	originData, err := json.Marshal(genesisData)
	genesisData.PermissionList = list
	genesisData.Version += 1
	newData, err := json.Marshal(genesisData)

	if err != nil {
		resp = respWrite(-1, "Phrase data to json failed.")
		logrus.Error(err)
		return
	}

	random := chameleonSec.ReSignature(elliptic.P256(), originData, newData, originRandom)
	genesisData.Id = hexKey
	genesisData.Random = hex.EncodeToString(random)

	tx, err := ethereum.CallContract(ethereum.MODIFY, originKey, hex.EncodeToString(random))

	if err != nil {
		resp = respWrite(-1, "Call ethereum contract failed.")
		logrus.Error(resp)
		return
	}

	genesisData.CreatedTx = tx
	rev, err := db.Put(context.TODO(), hexKey, genesisData)
	storageDataJson, err := json.Marshal(genesisData)

	if err != nil {
		resp = respWrite(-1, "Storage data serialize failed.")
		logrus.Error(resp)
		return
	}

	err = node.PutDataToNetwork(genesisData.Id, storageDataJson)

	if err != nil {
		resp = respWrite(-1, "Put data to P2P network failed.")
		logrus.Error(err)
		return
	}

	resp = respWrite(0, rev)
	return
}

func subPermissionRevoke(params *map[string]*structpb.Value) (resp *CommandRespond) {
	address := (*params)["address"].GetStringValue()

	genesisInt64, err := ethereum.ReadGenesis()
	genesisIndex := hex.EncodeToString(genesisInt64.Bytes())
	if err != nil {
		resp = respWrite(-1, "Read genesis index from contract failed.")
		logrus.Error(err)
		return
	}

	config := utils.ConfigInstance()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	chameleonSec := ecdcs.PrivateKey{}
	err = chameleonSec.FromHexString(hexPrivateKey)

	if err != nil {
		resp = respWrite(-1, "Read config file failed.")
		logrus.Error(err)
		return
	}

	dbUtil, err := utils.DBInstance()

	if err != nil {
		resp = respWrite(-1, "Get database instance failed.")
		logrus.Error(err)
		return
	}

	db := dbUtil.DB("data")

	var genesisData utils.GenesisData
	row := db.Get(context.TODO(), genesisIndex)
	err = row.ScanDoc(&genesisData)

	list := genesisData.PermissionList

	if !addressInList(&list, address) {
		resp = respWrite(-1, "Address not in list.")
		return
	}

	list = removeAddressFromList(&list, address)
	hexKey := genesisData.Id
	hexRandom := genesisData.Random

	originRandom, err := hex.DecodeString(hexRandom)

	if err != nil {
		resp = respWrite(-1, "Decode chameleon  data error.")
		logrus.Error(err)
		return
	}

	originKey, err := hex.DecodeString(hexKey)
	if err != nil {
		resp = respWrite(-1, "Decode chameleon  data error.")
		logrus.Error(err)
		return
	}

	genesisData.Id = ""
	genesisData.Random = ""
	genesisData.CreatedTx = ""

	originData, err := json.Marshal(genesisData)
	genesisData.PermissionList = list
	genesisData.Version += 1
	newData, err := json.Marshal(genesisData)

	if err != nil {
		resp = respWrite(-1, "Phrase data to json failed.")
		logrus.Error(err)
		return
	}

	random := chameleonSec.ReSignature(elliptic.P256(), originData, newData, originRandom)
	genesisData.Id = hexKey
	genesisData.Random = hex.EncodeToString(random)

	tx, err := ethereum.CallContract(ethereum.MODIFY, originKey, hex.EncodeToString(random))

	if err != nil {
		resp = respWrite(-1, "Call ethereum contract failed.")
		logrus.Error(resp)
		return
	}

	genesisData.CreatedTx = tx
	rev, err := db.Put(context.TODO(), hexKey, genesisData)
	storageDataJson, err := json.Marshal(genesisData)

	if err != nil {
		resp = respWrite(-1, "Storage data serialize failed.")
		logrus.Error(resp)
		return
	}

	err = node.PutDataToNetwork(genesisData.Id, storageDataJson)

	if err != nil {
		resp = respWrite(-1, "Put data to P2P network failed.")
		logrus.Error(err)
		return
	}

	resp = respWrite(0, rev)
	return
}

func addressInList(list *[]string, address string) bool {
	for _, item := range *list {
		if item == address {
			return true
		}
	}
	return false
}

func removeAddressFromList(list *[]string, address string) []string {
	for idx, item := range *list {
		if item == address {
			return append((*list)[:idx], (*list)[idx+1:]...)
		}
	}

	return *list
}

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
	"log"
	"time"
)

func GenesisProxy() (resp *CommandRespond) {
	// todo: 需要添加一个权限鉴别, 检查是否创建过创世数据以及存在权限

	config := utils.ConfigInstance()
	localAddress := config.Section("wallet").Key("address").String()
	hexPrivateKey := config.Section("wallet").Key("private").String()
	timestamp := time.Now().Unix()

	chameleonSec := ecdcs.PrivateKey{}
	err := chameleonSec.FromHexString(hexPrivateKey)
	chameleonPub, err := chameleonSec.ExportPublicKey(elliptic.P256())

	if err != nil {
		resp = respWrite(-1, "Parse private key error.")
		return
	}

	blockHash, err := ethereum.GetLatestBlockHash()
	var ret int32 = 0
	if err != nil {
		resp = respWrite(-1, "Get latest block hash error.")
		return
	}

	genesisData := utils.GenesisData{
		Id:             "",
		Random:         "",
		CreatedTx:      "",
		Version:        0,
		Operator:       localAddress,
		Timestamp:      timestamp,
		PermissionList: []string{},
		BlockHash:      blockHash,
		PublicKey:      chameleonPub.ToHexString(),
	}

	genesisJson, err := json.Marshal(genesisData)

	if err != nil {
		resp = respWrite(-1, "Genesis data serialization failed.")
		return
	}

	random, signature, err := chameleonSec.Signature(elliptic.P256(), genesisJson)
	if err != nil {
		resp = respWrite(-1, "Chameleon signature failed.")
		return
	}

	genesisData.Random = hex.EncodeToString(random)
	genesisData.Id = hex.EncodeToString(signature)

	// todo: 这里需要把bootstrap的信息通过配置文件获取
	tx, err := ethereum.ContractInitial(signature, hex.EncodeToString(random),
		"/ip4/testnet.decision01.com/tcp/5678")

	genesisData.CreatedTx = tx
	genesisJson, err = json.Marshal(genesisData)

	dbUtil, err := utils.DBInstance()

	if err != nil {
		resp = respWrite(-1, "Get database instance failed.")
		return
	}

	err = dbUtil.CreateDB(context.Background(), "data")

	if err != nil {
		resp = respWrite(-1, "Create table failed")
		log.Fatal(err)
	}

	db := dbUtil.DB("data")

	rev, err := db.Put(context.Background(), genesisData.Id, genesisData)
	if err != nil {
		resp = respWrite(-1, "Write database error.")
		log.Fatal(err)
	}

	dht, err := node.GetDHTInstance()

	err = node.PutDataToNetwork(genesisData.Id, genesisJson)
	if err != nil {
		resp = respWrite(-1, "Put data to P2P network failed.")
		log.Fatal(err)
		return
	}

	if err != nil {
		resp = respWrite(-1, "Peer node error.")
		return
	}

	dht.PutValue(context.Background(), genesisData.Id, genesisJson)

	resp = respWrite(0, rev)
	return
}

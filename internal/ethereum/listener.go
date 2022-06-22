package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/decision2016/osc/internal/node"
	"github.com/decision2016/osc/internal/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"log"
)

// ListenEthereumContract 监听以太坊合约的函数入口， 启一个goroutine来监听
func ListenEthereumContract() {
	cfg := utils.ConfigInstance()

	// 从配置文件中读取ws地址和需要监听的合约地址
	wssUrl := cfg.Section("contract").Key("ws").String()
	address := cfg.Section("contract").Key("address").String()

	client, err := ethclient.Dial(wssUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	// 使用配置文件中的地址来过滤日志信息
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
		return
	}

	logrus.Info("Start listen contract...")

	instance, err := NewEthereum(contractAddress, client)

	if err != nil {
		logrus.Error(err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// 监听到以太坊合约事件后注入的逻辑

			created, err := instance.ParseDataCreated(vLog)

			if err == nil {
				key := created.Key
				hexKey := hex.EncodeToString(key.Bytes())
				value, err := node.GetDataFromNetwork(hexKey)
				if err != nil {
					logrus.Error(err)
					continue
				}
				random := created.Random
				err = parseAndSaveData(value, hexKey, random)
				if err != nil {
					logrus.Error(err)
				}
				continue
			}

			modified, err := instance.ParseDataModified(vLog)
			if err == nil {
				key := modified.Key
				hexKey := hex.EncodeToString(key.Bytes())
				value, err := node.GetDataFromNetwork(hexKey)
				if err != nil {
					logrus.Error(err)
					continue
				}
				random := modified.Random
				err = parseAndSaveData(value, hexKey, random)
				if err != nil {
					logrus.Error(err)
				}
				continue
			}

			revoked, err := instance.ParseDataRevoked(vLog)
			if err == nil {
				logrus.Trace(revoked.Key)
				continue
			}
		}
	}
}

func parseAndSaveData(data []byte, key, random string) (err error) {
	var genesisData, dbGenesisData utils.GenesisData
	var storageData, dbStorageData utils.StorageData
	isGenesis := false

	err = json.Unmarshal(data, &storageData)

	if err != nil {
		isGenesis = true
		err := json.Unmarshal(data, &genesisData)

		if err != nil {
			return
		}
	}

	dbUtil, err := utils.DBInstance()

	if err != nil {
		logrus.Error(err)
		return
	}

	db := dbUtil.DB("data")

	if isGenesis {
		row := db.Get(context.TODO(), key)
		err := row.ScanDoc(&dbGenesisData)
		if err != nil {
			logrus.Error(err)
			return err
		}

		if dbGenesisData.Id == "" || dbGenesisData.Version < genesisData.Version {
			_, err := db.Put(context.TODO(), key, genesisData)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
	} else {
		row := db.Get(context.TODO(), key)
		err := row.ScanDoc(&dbStorageData)
		if err != nil {
			logrus.Error(err)
			return err
		}

		if dbStorageData.Id == "" || dbStorageData.Version < genesisData.Version {
			_, err := db.Put(context.TODO(), key, storageData)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
	}
	return
}

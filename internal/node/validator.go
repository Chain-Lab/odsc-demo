package node

import (
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"github.com/decision2016/go-crypto/ecdcs"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
)

type ChameleonValidator struct{}

func (v ChameleonValidator) Validate(key string, value []byte) (err error) {
	return nil

	storageData := utils.StorageData{}
	genesisData := utils.GenesisData{}
	isGenesisData := false
	err = json.Unmarshal(value, &storageData)
	logrus.Trace(err)

	if err != nil {
		logrus.Info("Decode as storage data failed.")
		err = json.Unmarshal(value, &genesisData)
		logrus.Info(err)

		if err != nil {
			return
		}

		isGenesisData = true
	}

	copyGenesisData := genesisData
	copyStorageData := storageData
	var jsonData []byte
	var random string

	var hexPublicKey string
	if isGenesisData {
		hexPublicKey = genesisData.PublicKey
		copyGenesisData.CreatedTx = ""
		copyGenesisData.Random = ""
		copyGenesisData.Id = ""
	} else {
		hexPublicKey = storageData.PublicKey
		copyStorageData.CreatedTx = ""
		copyStorageData.Random = ""
		copyStorageData.Id = ""
	}
	logrus.Trace(genesisData.PublicKey)

	chameleonPub := ecdcs.PublicKey{}
	err = chameleonPub.FromHexString(hexPublicKey)

	if err != nil {
		return
	}

	if isGenesisData {
		jsonData, err = json.Marshal(copyGenesisData)
		if err != nil {
			return
		}
		random = genesisData.Random

	} else {
		jsonData, err = json.Marshal(copyStorageData)
		if err != nil {
			return
		}
		random = storageData.Random
	}

	logrus.Info(chameleonPub.ToHexString())
	result, err := chameleonPub.Verify(elliptic.P256(), jsonData, key, random)

	if !result {
		return errors.New("verify chameleon failed")
	}

	return
}

func (v ChameleonValidator) Select(key string, values [][]byte) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("can't select from no values")
	}

	return 0, nil
}

package node

import (
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"github.com/decision2016/go-crypto/ecdcs"
	"github.com/decision2016/osc/internal/utils"
)

type ChameleonValidator struct{}

func (v ChameleonValidator) Validate(key string, value []byte) (err error) {
	return nil

	storageData := utils.StorageData{}
	genesisData := utils.GenesisData{}
	isGenesisData := false
	err = json.Unmarshal(value, &storageData)

	if err != nil {
		err = json.Unmarshal(value, &genesisData)

		if err != nil {
			return
		}

		isGenesisData = true
	}

	copyGenesisData := genesisData
	copyStorageData := storageData
	var jsonData []byte
	var signature string
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
		signature, random = genesisData.Id, genesisData.Random

	} else {
		jsonData, err = json.Marshal(copyStorageData)
		if err != nil {
			return
		}
		signature, random = storageData.Id, storageData.Random
	}

	result, err := chameleonPub.Verify(elliptic.P256(), jsonData, signature, random)

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

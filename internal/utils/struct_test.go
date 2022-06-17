package utils

import (
	"crypto/elliptic"
	"encoding/json"
	"github.com/decision2016/go-crypto/ecdcs"
	"testing"
)

func TestStruct(t *testing.T) {
	hexPrivate := "3f707e43bd8805f379b86afacc70f39767e9a3ca896553445d30d8c5ad089610"
	chameleonSec := ecdcs.PrivateKey{}
	chameleonSec.FromHexString(hexPrivate)

	chameleonPub, _ := chameleonSec.ExportPublicKey(elliptic.P256())

	// 验证不同struct在反序列化的时候是否会出现错误
	genesisData := GenesisData{
		Id:             "",
		Random:         "",
		CreatedTx:      "",
		Version:        0,
		Operator:       "0x9a8489051f6fA66B3104920844B1ccc7C3F8eDbC",
		Timestamp:      1655298330,
		PermissionList: []string{"0x9a8489051f6fA66B3104920844B1ccc7C3F8eDbC"},
		BlockHash:      "0xb2658061b110ef6a22503d65464c990dbf948c7ba2e1f401157359a7fb1bfa50",
		PublicKey:      chameleonPub.ToHexString(),
	}

	genesisBytes, _ := json.Marshal(genesisData)
	t.Log(string(genesisBytes))
	storageData := StorageData{}

	err := json.Unmarshal(genesisBytes, &storageData)
	t.Log(err)
	t.Log(storageData.Data)
}

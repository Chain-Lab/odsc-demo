package utils

import (
	"encoding/json"
	"testing"
)

func TestStruct(t *testing.T) {
	// 验证不同struct在反序列化的时候是否会出现错误
	genesisData := GenesisData{
		Id:             "test",
		PermissionList: []string{string([]byte("Test"))},
		PublicKey:      "test1",
	}

	genesisBytes, _ := json.Marshal(genesisData)
	t.Log(string(genesisBytes))
	storageData := StorageData{}

	err := json.Unmarshal(genesisBytes, &storageData)
	t.Log(err)
	t.Log(storageData.Data)
}

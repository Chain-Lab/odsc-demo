package utils

// structs for json encoding

type StorageData struct {
	Id        string `json:"id,omitempty"`
	Version   int32  `json:"version,omitempty"`
	Random    string `json:"random,omitempty"`
	CreatedTx string `json:"created_tx,omitempty"`
	Operator  string `json:"operator,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      []byte `json:"data,omitempty"`
	BlockHash string `json:"block_hash,omitempty"`
}

type GenesisData struct {
	Id             string   `json:"id,omitempty"`         // 变色龙哈希结果 (哈希后写入)
	Version        int32    `json:"version,omitempty"`    // 数据更新的版本号 (哈希前写入)
	Random         string   `json:"random,omitempty"`     // 数据哈希的随机数 (哈希后写入)
	CreatedTx      string   `json:"created_tx,omitempty"` // 数据创建的交易哈希 (哈希后写入)
	Operator       string   `json:"operator,omitempty"`   // 权限管理者地址 (哈希前写入)
	PublicKey      string   `json:"public_key,omitempty"` // 数据创建者的公钥信息
	Timestamp      int64    `json:"timestamp,omitempty"`  // 创建的时间戳 (哈希前写入)
	PermissionList []string `json:"data,omitempty"`       // 权限表 (哈希前写入)
	BlockHash      string   `json:"block_hash,omitempty"` // 当前的区块哈希(哈希前写入)
}

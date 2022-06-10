package utils

// structs for json encoding

type StorageData struct {
	Id        string `json:"id,omitempty"`
	Version   int32  `json:"version,omitempty"`
	Random    string `json:"random,omitempty"`
	CreatedTx string `json:"created_tx,omitempty"`
	Operator  string `json:"public_key,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      string `json:"data,omitempty"`
	BlockHash string `json:"block_hash,omitempty"`
}

type GenesisData struct {
	Id             string `json:"id,omitempty"`
	Version        int32  `json:"version,omitempty"`
	Random         string `json:"random,omitempty"`
	CreatedTx      string `json:"created_tx,omitempty"`
	Operator       string `json:"operator,omitempty"`
	Timestamp      int64  `json:"timestamp,omitempty"`
	PermissionList string `json:"permission_list,omitempty"`
	BlockHash      string `json:"block_hash,omitempty"`
}

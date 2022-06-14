package ethereum

import (
	"log"
	"testing"
)

func TestGetBootstrap(t *testing.T) {
	genesis, _ := ReadGenesis()

	log.Println(genesis.String() == "0")
}

func TestGetLatestBlock(t *testing.T) {
	hash, _ := GetLatestBlockHash()

	log.Println(hash)
}
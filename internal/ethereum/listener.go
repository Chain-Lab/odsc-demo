package ethereum

import (
	"context"
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

	log.Println("Start listen contract...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// 监听到以太坊合约事件后注入的逻辑

			fmt.Println(vLog)
		}
	}
}

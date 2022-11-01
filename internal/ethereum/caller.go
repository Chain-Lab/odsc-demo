package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/decision2016/osc/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type CallType int

const (
	CREATE CallType = 1
	MODIFY CallType = 2
	REVOKE CallType = 3
)

func ReadBootstrap() string {
	config := utils.ConfigInstance()
	contractAddress := config.Section("contract").Key("address").String()
	httpUrl := config.Section("contract").Key("http").String()

	client, err := ethclient.Dial(httpUrl)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewEthereum(address, client)

	if err != nil {
		log.Fatal(err)
	}

	bootstrap, err := instance.Bootstrap(nil)

	if err != nil {
		log.Fatal(err)
	}

	return bootstrap
}

func ContractInitial(genesisKey []byte, random string, bootstrap string) (tx string, err error) {
	config := utils.ConfigInstance()
	contractAddress := config.Section("contract").Key("address").String()
	httpUrl := config.Section("contract").Key("http").String()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	// 从配置文件中加载私钥并且连接到以太坊测试网络
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	client, err := ethclient.Dial(httpUrl)

	if err != nil {
		log.Fatal(err)
	}

	auth, err := buildAuth(client, privateKey)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewEthereum(address, client)

	genesis := big.Int{}
	genesis.SetBytes(genesisKey)

	transaction, err := instance.Initialization(auth, &genesis, []byte(random), bootstrap)

	if err != nil {
		log.Fatal(err)
	}

	tx = transaction.Hash().String()

	return
}

func CallContract(funcType CallType, key []byte, random string) (tx string, err error) {
	config := utils.ConfigInstance()
	contractAddress := config.Section("contract").Key("address").String()
	httpUrl := config.Section("contract").Key("http").String()
	hexPrivateKey := config.Section("wallet").Key("private").String()

	// 从配置文件中加载私钥并且连接到以太坊测试网络
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	client, err := ethclient.Dial(httpUrl)

	if err != nil {
		log.Fatal(err)
	}

	auth, err := buildAuth(client, privateKey)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewEthereum(address, client)

	keyUint := big.Int{}
	keyUint.SetBytes(key)

	switch funcType {
	case CREATE:
		{
			transaction, err := instance.Create(auth, &keyUint, []byte(random))
			if err != nil {
				log.Fatal(err)
			}
			tx = transaction.Hash().String()
		}
	case MODIFY:
		{
			transaction, err := instance.Modify(auth, &keyUint, []byte(random))
			if err != nil {
				log.Fatal(err)
			}
			tx = transaction.Hash().String()
		}
	case REVOKE:
		{
			transaction, err := instance.Revoke(auth, &keyUint)
			if err != nil {
				log.Fatal(err)
			}
			tx = transaction.Hash().String()
		}
	}

	return
}

func ReadGenesis() (genesis *big.Int, err error) {
	config := utils.ConfigInstance()
	contractAddress := config.Section("contract").Key("address").String()
	httpUrl := config.Section("contract").Key("http").String()

	client, err := ethclient.Dial(httpUrl)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewEthereum(address, client)

	if err != nil {
		log.Fatal(err)
	}

	genesis, err = instance.Genesis(nil)

	if err != nil {
		log.Fatal(err)
	}

	return
}

func buildAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (auth *bind.TransactOpts, err error) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取到当前地址的nonce信息
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	chainId, err := client.ChainID(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return
}

func GetLatestBlockHash() (blockHash string, err error) {
	config := utils.ConfigInstance()
	httpUrl := config.Section("contract").Key("http").String()

	client, err := ethclient.Dial(httpUrl)

	header, err := client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	blockHash = header.Hash().String()
	return
}

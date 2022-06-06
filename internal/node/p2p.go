package node

import (
	"context"
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"github.com/libp2p/go-libp2p"
	host2 "github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-kad-dht"
	"log"
)

var config = utils.ConfigInstance()

var (
	host        host2.Host
	ctx                      = context.Background()
	kademliaDHT *dht.IpfsDHT = nil
)

// start 启动P2P节点， 仅运行一次作为DHT的初始化
// 在后续需要进行P2P网络的操作的时候获取一个实例来进行操作
func start() (err error) {
	runAsBootstrap, err := config.Section("node").Key("run_as_bootstrap").Bool()

	if err != nil {
		log.Fatal("Get config params error.")
		return
	}

	if runAsBootstrap {
		err := runBootstrap()
		if err != nil {
			log.Fatal("Start bootstrap service failed.")
			return
		}
	} else {
		err := runPeerNode()
		if err != nil {
			log.Fatal("Start peer node failed.")
			return
		}
	}

	kademliaDHT, err = dht.New(ctx, host)

	if err != nil {
		log.Fatal("Start kademliaDHT service failed.")
		return
	}

	return
}

func GetDHTInstance() (instance *dht.IpfsDHT, err error) {
	if kademliaDHT == nil {
		err := start()
		if err != nil {
			log.Fatal("DHT initialization failed.")
			return nil, err
		}
	}

	return kademliaDHT, err
}

func runBootstrap() (err error) {
	config = utils.ConfigInstance()
	port, err := config.Section("node").Key("port").Int()

	if err != nil {
		return
	}

	listenAddr := fmt.Sprintf("ip4/0.0.0.0/tcp/%d", port)

	host, err = libp2p.New(
		libp2p.ListenAddrStrings(
			listenAddr,
		),
		libp2p.DefaultTransports,
	)

	return
}

func runPeerNode() (err error) {

	return
}

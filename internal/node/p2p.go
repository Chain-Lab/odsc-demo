package node

import (
	"context"
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"github.com/libp2p/go-libp2p"
	host2 "github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var config = utils.ConfigInstance()

var (
	host        host2.Host
	ctx                      = context.Background()
	kademliaDHT *dht.IpfsDHT = nil
)

var p2pOnce sync.Once

// Start 启动P2P节点， 仅运行一次作为DHT的初始化
// 在后续需要进行P2P网络的操作的时候获取一个实例来进行操作
func Start() (err error) {
	runAsBootstrap, err := config.Section("node").Key("run_as_bootstrap").Bool()

	if err != nil {
		logrus.Fatal("Get config params error.")
		return
	}

	err = runBootstrap()
	if err != nil {
		logrus.Fatal("Start bootstrap service failed: ", err)
		return
	}

	logrus.Info("Bootstrap node start.")

	if !runAsBootstrap {
		err := runPeerNode()
		if err != nil {
			logrus.Fatal("Start peer node failed.")
		}
	}

	kademliaDHT, err = dht.New(ctx, host, dht.Validator(ChameleonValidator{}), dht.ProtocolPrefix("/chameleon"))

	if err != nil {
		logrus.Info(err)
		logrus.Fatal("Start kademliaDHT service failed.")
		return
	}

	go func() {
		for {
			logrus.Info(kademliaDHT.RoutingTable())
			time.Sleep(10 * time.Second)
		}
	}()

	return
}

func GetDHTInstance() (instance *dht.IpfsDHT, err error) {
	logrus.Trace("Start DHT initial.")
	p2pOnce.Do(func() {
		err := Start()
		if err != nil {
			logrus.Fatal("DHT initialization failed.")
		}
	})

	return kademliaDHT, err
}

func runBootstrap() (err error) {
	config = utils.ConfigInstance()
	port, err := config.Section("node").Key("port").Int()

	if err != nil {
		logrus.Fatal(err)
		return
	}

	listenAddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)

	host, err = libp2p.New(
		libp2p.ListenAddrStrings(
			listenAddr,
		),
		libp2p.DefaultTransports,
	)

	logrus.Info("Local node: ", host.Addrs())
	logrus.Info("Node Id: ", host.ID())

	return
}

func runPeerNode() (err error) {
	config = utils.ConfigInstance()
	bootstrap := config.Section("node").Key("bootstrap").String()

	addr, err := multiaddr.NewMultiaddr(bootstrap)
	if err != nil {
		logrus.Fatal(err)
	}
	peer, err := peerstore.AddrInfoFromP2pAddr(addr)
	if err != nil {
		logrus.Fatal(err)
	}
	err = host.Connect(context.Background(), *peer)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connect to P2P network.")

	return
}

func PutDataToNetwork(key string, value []byte) (err error) {
	return kademliaDHT.PutValue(context.Background(), key, value, dht.Quorum(1))
}

func GetDataFromNetwork(key string) (value []byte, err error) {
	return kademliaDHT.GetValue(context.Background(), key)
}

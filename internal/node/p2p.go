package node

import (
	"context"
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"github.com/libp2p/go-libp2p"
	host2 "github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
	"time"
)

var config = utils.ConfigInstance()

var (
	host             host2.Host
	ctx                                        = context.Background()
	kademliaDHT      *dht.IpfsDHT              = nil
	routingDiscovery *routing.RoutingDiscovery = nil
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

	if runAsBootstrap {
		kademliaDHT, err = dht.New(ctx, host,
			dht.Validator(ChameleonValidator{}),
			dht.ProtocolPrefix("/chameleon"),
			dht.Mode(dht.ModeServer))
	} else {
		kademliaDHT, err = dht.New(ctx, host,
			dht.Validator(ChameleonValidator{}),
			dht.ProtocolPrefix("/chameleon"))
	}

	if err != nil {
		logrus.Info(err)
		logrus.Fatal("Start kademliaDHT service failed.")
		return
	}

	err = kademliaDHT.Bootstrap(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	go discover(ctx, host)

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

// PutDataToNetwork 将数据存放到P2P网络中
// @key: key-value数据的键，通常使用变色龙哈希结果
// @value: key-value数据的值
func PutDataToNetwork(key string, value []byte) (err error) {
	// 将数据放入到P2P网络上
	return kademliaDHT.PutValue(ctx, key, value)
}

func GetDataFromNetwork(key string) (value []byte, err error) {
	return kademliaDHT.GetValue(ctx, key)
}

func discover(ctx context.Context, h host2.Host) (err error) {
	if routingDiscovery == nil {
		routingDiscovery = routing.NewRoutingDiscovery(kademliaDHT)
	}

	// 调用方法加载路由表， 发现网络中的其他节点
	_, err = routingDiscovery.Advertise(ctx, "test111")
	logrus.Info("Get routing discovery instance...")

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logrus.Info("Context done.")
			return
		case <-ticker.C:
			peers, err := routingDiscovery.FindPeers(ctx, "test111")
			if err != nil {
				logrus.Fatal(err)
			}

			for p := range peers {
				// logrus.Info(p.String())
				if p.ID == h.ID() {
					continue
				}
				if h.Network().Connectedness(p.ID) != network.Connected {
					logrus.Info("Connect to node ", p.ID)
					_, err = h.Network().DialPeer(ctx, p.ID)
					if err != nil {
						logrus.Error(err)
						continue
					}
				}
			}
		}
	}
}

package rpc

import (
	"fmt"
	"github.com/decision2016/osc/internal/utils"
	"log"
	"net"
	"net/rpc"
)

func StartRPCService() (err error) {
	config := utils.ConfigInstance()
	port, err := config.Section("rpc").Key("port").Int()

	if err != nil {
		return
	}

	rpc.RegisterName("CommandService", new(CommandService))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal("Listen TCP error: ", err)
		return
	}

	conn, err := listener.Accept()

	if err != nil {
		log.Fatal("Accept error: ", err)
		return
	}

	rpc.ServeConn(conn)

	return
}

package domain

import (
	"fmt"
	"net"

	respv2 "github.com/codecrafters-io/redis-starter-go/internal/resp_v2"
)

type ValidHandlers string

const (
	Ping ValidHandlers = "ping"
	Echo ValidHandlers = "echo"
	Set  ValidHandlers = "set"
	Get  ValidHandlers = "get"
)

type Params struct {
	C    net.Conn
	Data any
}

var DomainHandlers = map[ValidHandlers]func(Params){
	Ping: PingHandler,
	Echo: EchoHandler,
	Set:  SetHandler,
	Get:  GetHandler,
}

func ApplyHandler(rootNode respv2.Node, c net.Conn) {
	arrayNode, ok := rootNode.(respv2.ArrayNode)

	if !ok {
		fmt.Errorf("failed to get bulkString handler!")
		return
	}

	headerNode := arrayNode.Data[0]

	switch h := headerNode.(type) {
	case respv2.BulkStringNode:
		handleBulkString(h)
	}
}

func handleBulkString(nodes respv2.BulkStringNode) {
	// if nodes.Data

}

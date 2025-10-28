package domain

import "net"

type ValidHandlers string

const (
	Ping ValidHandlers = "ping"
	Pong ValidHandlers = "pong"
	Echo ValidHandlers = "echo"
)

var DomainHandlers = map[ValidHandlers]func(net.Conn){
	Ping: PingHandler,
}

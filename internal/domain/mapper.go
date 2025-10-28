package domain

import "net"

type ValidHandlers string

const (
	Ping ValidHandlers = "ping"
	Echo ValidHandlers = "echo"
)

type Params struct {
	C    net.Conn
	Data any
}

var DomainHandlers = map[ValidHandlers]func(Params){
	Ping: PingHandler,
	Echo: EchoHandler,
}

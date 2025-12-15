package domain

import (
	"fmt"
	"net"
	"slices"
	"strings"
)

type ValidHandlers string

const (
	Ping  ValidHandlers = "ping"
	Echo  ValidHandlers = "echo"
	Debug ValidHandlers = "debug"
	Get   ValidHandlers = "get"
	Set   ValidHandlers = "set"
)

type Params struct {
	C    net.Conn
	Data any
}

var DomainHandlers = map[ValidHandlers]func(Params){
	Ping: PingHandler,
	// Echo:  EchoHandler,
	Debug: DebugHandler,
	// Get:   GetHandler,
	// Set: SetHandler,
}

func IsValidOp(op string) bool {
	return slices.Contains([]ValidHandlers{Ping, Echo, Debug, Get, Set}, ValidHandlers(op))
}

func GetHandlerForRequestedOperation(op string) (func(Params), error) {
	domainOperation := ValidHandlers(strings.ToLower(op))
	domainHandler := DomainHandlers[domainOperation]
	if domainHandler == nil {
		return nil, fmt.Errorf("failed to match op %s", op)
	}
	return domainHandler, nil
}

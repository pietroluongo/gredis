package domain

import (
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
)

func PingHandler(c net.Conn) {
	c.Write([]byte(output.BuildSimpleString("PONG")))
}

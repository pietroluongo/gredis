package domain

import (
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
)

func PingHandler(p Params) {
	p.C.Write([]byte(output.BuildSimpleString("PONG")))
}

func PingHandlerV2(p BaseParams) error {
	p.C.Write([]byte(output.BuildSimpleString("PONG")))
	return nil
}

package domain

import (
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

func PingHandler(p Params) {
	storage.Debug()
	p.C.Write([]byte(output.BuildSimpleString("PONG")))
}

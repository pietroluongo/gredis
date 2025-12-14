package domain

import (
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type SetParams struct {
	Params
	key   string
	value string
}

func (SetParams) params() {}

func SetHandlerV2(p SetParams) {
	storage.Set(p.key, p.value)
	p.C.Write([]byte(output.BuildSimpleString("OK")))
}

func SetHandler(p Params) {
	// args := p.Data.([]resp.RespMessage)
	storage.Set(args[0].Content.(string), args[1].Content.(string))
	p.C.Write([]byte(output.BuildSimpleString("OK")))
}

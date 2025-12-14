package domain

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

func GetHandler(p Params) {
	args := p.Data.([]resp.RespMessage)
	fmt.Printf("%v", p.Data)
	storage.Get(args[1].Content.(string))
	p.C.Write([]byte(output.BuildSimpleString("OK")))
}

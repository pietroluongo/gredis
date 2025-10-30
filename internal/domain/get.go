package domain

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

func GetHandler(p Params) {
	params := p.Data.([]resp.RespMessage)
	fmt.Printf("Requested %s", params[0].Content)
	result := storage.Get(params[0].Content.(string))
	p.C.Write([]byte(output.BuildBulkString(result)))
}

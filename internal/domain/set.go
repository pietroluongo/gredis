package domain

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

func SetHandler(p Params) {
	data := p.Data.([]resp.RespMessage)

	key := data[0].Content.(string)
	val := data[1].Content.(string)
	fmt.Printf("Setting %s to %s\n", key, val)
	storage.Set(key, val)
	p.C.Write([]byte("+OK\r\n"))
}

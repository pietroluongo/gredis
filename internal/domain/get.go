package domain

import (
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

func GetHandler(p Params) {
	result := storage.Get(p.Data.(string))
	p.C.Write([]byte(result))
}

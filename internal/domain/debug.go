package domain

import "github.com/codecrafters-io/redis-starter-go/internal/storage"

func DebugHandler(p Params) {
	storage.Debug()
}

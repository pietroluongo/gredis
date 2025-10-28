package domain

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
)

func EchoHandler(p Params) {
	args := p.Data.([]resp.RespMessage)
	var response = []string{}
	for _, result := range args {
		if result.Kind != resp.BulkString && result.Kind != resp.SimpleString {
			fmt.Printf("WARN: trying to echo something that is not echoable")
			return
		}
		response = append(response, result.Content.(string))
	}
	p.C.Write([]byte(output.BuildSimpleString(strings.Join(response, " "))))
}

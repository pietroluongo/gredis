package domain

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/respv2"
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

type EchoParams struct {
	BaseParams BaseParams
	Args       []respv2.BulkStringNode
}

func (EchoParams) param() {}

func BuildEchoParams(p BuilderFunctionParams) ParamsV2 {
	return EchoParams{BaseParams: BaseParams{C: p.C}, Args: p.Args}
}

func EchoHandlerV2(p EchoParams) error {
	var response = []string{}
	for _, r := range p.Args {
		response = append(response, r.Data)
	}
	p.BaseParams.C.Write([]byte(output.BuildBulkString(strings.Join(response, " "))))
	return nil
}

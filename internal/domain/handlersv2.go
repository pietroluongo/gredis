package domain

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/respv2"
)

type BuilderFunctionParams struct {
	C    net.Conn
	Args []respv2.BulkStringNode
}

type ParamsV2 interface {
	param()
}

type BaseParams struct {
	C net.Conn
}

func (BaseParams) param() {}

func buildBaseParams(p BuilderFunctionParams) ParamsV2 {
	return BaseParams{C: p.C}
}

type HandlerFunc[T ParamsV2] func(T) error

func wrapHandler[T ParamsV2](h HandlerFunc[T]) func(ParamsV2) error {
	return func(p ParamsV2) error {
		typedParam, ok := p.(T)
		if !ok {
			return fmt.Errorf("failed to get typed param, expected %T, got %T", *new(T), p)
		}
		return h(typedParam)
	}
}

type HandlerObject struct {
	Handler func(ParamsV2) error
	Builder func(BuilderFunctionParams) ParamsV2
}

var DomainHandlersV2 = map[ValidHandlers]*HandlerObject{
	Echo: {Handler: wrapHandler(EchoHandlerV2), Builder: BuildEchoParams},
	Ping: {Handler: wrapHandler(PingHandlerV2), Builder: buildBaseParams},
}

func GetHandlerAndBuilder(op []respv2.BulkStringNode) (func(ParamsV2) error, func(BuilderFunctionParams) ParamsV2, error) {
	domainHandler := DomainHandlersV2[ValidHandlers(strings.ToLower(op[0].Data))]

	if domainHandler == nil {
		return nil, nil, fmt.Errorf("failed to match op %s to any handler", op)
	}
	return domainHandler.Handler, domainHandler.Builder, nil
}

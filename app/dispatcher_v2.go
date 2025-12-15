package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/domain"
	"github.com/codecrafters-io/redis-starter-go/internal/logger"
	"github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	"github.com/codecrafters-io/redis-starter-go/internal/respv2"
)

func dispatchV2(message respv2.CommandNode, c net.Conn) {
	logger.Log.Info(fmt.Sprintf("dispatch v2 called %v", message), "message", message)
	handler, builder, err := domain.GetHandlerAndBuilder(message.Command)
	if err != nil {
		logger.Log.Error("Failed to match handler", "error", err)
		c.Write([]byte(output.BuildSimpleError(fmt.Sprintf("Could not find command %s", message.Command[0].Data))))
		return
	}

	builtParams := builder(domain.BuilderFunctionParams{C: c, Args: message.Args})
	handlerErr := handler(builtParams)
	logger.Log.Info("success processing command", "command", message, "client", c.RemoteAddr(), "err", handlerErr)
}

package main

import (
	"fmt"
	"net"
	"slices"
	"strings"

	domain "github.com/codecrafters-io/redis-starter-go/internal/domain"
	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	respOutput "github.com/codecrafters-io/redis-starter-go/internal/resp/output"
)

func convertMessageToMessageArray(m resp.RespMessage) resp.ArrayRespMessage {
	return resp.ArrayRespMessage{Kind: resp.Array, Content: m.Content.([]resp.RespMessage)}
}

func dispatch(message resp.RespMessage, c net.Conn) {
	log.Info("dispatch called")
	if message.Kind == resp.Array {
		log.Info("is array")
		msgArr := convertMessageToMessageArray(message)
		if checkIfArrayIsCommand(msgArr) {
			log.Info("array is command")
			dispatchCommandWithArray(msgArr, c)
			return
		}
	}
	log.Info("is not array")

	if message.Kind == resp.SimpleString && isOperation(strings.ToLower(message.Content.(string))) {
		log.Info("is simple op")
		dispatchSimpleOperation(message, c)
		return
	}

}

func checkIfArrayIsCommand(r resp.ArrayRespMessage) bool {
	possibleCommandItem := r.Content[0]
	log.Info(fmt.Sprintf("possible command item is %v", possibleCommandItem))
	return slices.Contains([]RedisCommands{ping, echo}, RedisCommands(strings.ToLower(possibleCommandItem.Content.(string))))
}

func dispatchCommandWithArray(r resp.ArrayRespMessage, c net.Conn) {
	commandRequest := r.Content[0]
	domainOperation := domain.ValidHandlers(strings.ToLower(commandRequest.Content.(string)))
	domainHandler := domain.DomainHandlers[domainOperation]

	if domainHandler == nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return
	}
	commandParams := r.Content[1:]
	domain.DomainHandlers[domainOperation](domain.Params{C: c, Data: commandParams})

}

func isOperation(s string) bool {
	return slices.Contains([]domain.ValidHandlers{domain.Ping, domain.Echo}, domain.ValidHandlers(s))
}

func dispatchSimpleOperation(m resp.RespMessage, c net.Conn) {
	domainOperation := domain.ValidHandlers(strings.ToLower(m.Content.(string)))
	domainHandler := domain.DomainHandlers[domainOperation]
	if domainHandler == nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return
	}
	domain.DomainHandlers[domainOperation](domain.Params{C: c})
}

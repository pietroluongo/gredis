package main

import (
	"fmt"
	"net"
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

	if message.Kind == resp.SimpleString && domain.IsValidOp(strings.ToLower(message.Content.(string))) {
		log.Info("is simple op")
		dispatchSimpleOperation(message, c)
		return
	}

}

func checkIfArrayIsCommand(r resp.ArrayRespMessage) bool {
	possibleCommandItem := r.Content[0]
	log.Info(fmt.Sprintf("possible command item is %v", possibleCommandItem))
	requestedOp := strings.ToLower(possibleCommandItem.Content.(string))
	return domain.IsValidOp(requestedOp)
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

func dispatchSimpleOperation(m resp.RespMessage, c net.Conn) {
	requestedOp := strings.ToLower(m.Content.(string))
	domainHandler, err := domain.GetHandlerForRequestedOperation(requestedOp)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s %s", requestedOp, err.Error()))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return
	}
	domainHandler(domain.Params{C: c})
}

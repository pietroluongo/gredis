package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/domain"
	respOutput "github.com/codecrafters-io/redis-starter-go/internal/resp/output"
	respv2 "github.com/codecrafters-io/redis-starter-go/internal/resp_v2"
)

func dispatch(message respv2.Node, c net.Conn) {
	log.Info("dispatch called")

	if msg, ok := message.(respv2.SimpleStringNode); ok {
		log.Info("is simple string")
		if msg.IsSimpleCommand() {
			dispatchSimpleOperation(msg, c)
			return
		}
	}

	if msg, ok := message.(respv2.CommandNode); ok {
		log.Info("is array")
		fmt.Printf("the command is %v\n", msg)
		dispatchArrayBasedCommand(msg, c)
	}

	// if message.IsArray() {
	// 	log.Info("is array")
	// 	msgArr := message.AsMessageArray()
	// 	if arrayIsCommand(msgArr) {
	// 		log.Info("array is command")
	// 		dispatchArrayBasedCommand(msgArr, c)
	// 		return
	// 	}
	// }
	// log.Info("is not array")

	// if message.Kind == resp.SimpleString && isSimpleCommandOperation(strings.ToLower(message.Content.(string))) {
	// 	log.Info("is simple op")
	// 	dispatchSimpleOperation(message, c)
	// 	return
	// }

}

func dispatchArrayBasedCommand(r respv2.CommandNode, c net.Conn) error {
	commandRequest := r.Command

	domainOperation := domain.ValidHandlers(strings.ToLower(commandRequest.Data))
	domainHandler := domain.DomainHandlers[domainOperation]

	if domainHandler == nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return fmt.Errorf("failed to match handler for op %s", domainOperation)
	}
	domain.DomainHandlers[domainOperation](domain.Params{C: c, Data: r.Args})
	return nil
}

// func dispatchSimpleOperation(m resp.RespMessage, c net.Conn) {
// 	domainOperation := domain.ValidHandlers(strings.ToLower(m.Content.(string)))
// 	domainHandler := domain.DomainHandlers[domainOperation]
// 	if domainHandler == nil {
// 		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
// 		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
// 		return
// 	}
// 	domain.DomainHandlers[domainOperation](domain.Params{C: c})
// }

func dispatchSimpleOperation(m respv2.SimpleStringNode, c net.Conn) {
	domainOperation := domain.ValidHandlers(strings.ToLower(m.Data))
	domainHandler := domain.DomainHandlers[domainOperation]

	if domainHandler == nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return
	}
	domain.DomainHandlers[domainOperation](domain.Params{C: c})
}

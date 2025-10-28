package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"slices"

	domain "github.com/codecrafters-io/redis-starter-go/internal/domain"
	resp "github.com/codecrafters-io/redis-starter-go/internal/resp"
	respOutput "github.com/codecrafters-io/redis-starter-go/internal/resp/output"
)

var log *slog.Logger

func initLog() {
	file, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE, 0644)

	var logOutput io.Writer

	if err != nil {
		slog.Default().Warn(fmt.Sprintf("Failed to create log file, using only stdout %s", err.Error()))
		logOutput = os.Stdout
	} else {
		logOutput = io.MultiWriter(os.Stdout, file)
	}
	log = slog.New(slog.NewTextHandler(logOutput, &slog.HandlerOptions{AddSource: false})).With(slog.Group("context", "package", "server"))
}

func server(connection net.Conn, parentChannel chan TCPStatus) {
	initLog()

	command := make([]byte, 1024)
	for {
		clear(command)
		size, err := connection.Read(command)
		if err != nil {
			log.Error(fmt.Sprint("Failed to read data from connection ", err.Error()))
			parentChannel <- TCPStatus{isError: true}
			return
		}
		log.Info(fmt.Sprintf("Read %d bytes", size))

		message := resp.ParseMessage(command[:size])

		log.Info(fmt.Sprintf("Got following message: %v", message))

		if message.Kind == resp.SimpleString && isOperation(message.Content.(string)) {
			log.Info("dispatching")
			dispatchOperation(message, connection)
			continue
		}

		connection.Write([]byte(respOutput.BuildSimpleString("OK")))
	}
}

func isOperation(s string) bool {
	return slices.Contains([]domain.ValidHandlers{domain.Ping, domain.Echo}, domain.ValidHandlers(s))
}

func dispatchOperation(message resp.RespMessage, c net.Conn) {
	domainOperation := domain.ValidHandlers(message.Content.(string))
	domainHandler := domain.DomainHandlers[domainOperation]
	if domainHandler == nil {
		log.Error(fmt.Sprintf("Failed to match handler for operation %s", domainOperation))
		c.Write([]byte(respOutput.BuildSimpleError("Matched operator, but failed to match handler")))
		return
	}
	domain.DomainHandlers[domainOperation](c)
}

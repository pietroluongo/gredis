package main

import (
	"fmt"
	"io"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/logger"
	"github.com/codecrafters-io/redis-starter-go/internal/respv2"
)

// var log *slog.Logger

// func initLog() {
// 	file, err := os.OpenFile("./log.json", os.O_RDWR|os.O_CREATE, 0644)

// 	var logOutput io.Writer

// 	if err != nil {
// 		slog.Default().Warn(fmt.Sprintf("Failed to create log file, using only stdout %s", err.Error()))
// 		logOutput = os.Stdout
// 	} else {
// 		logOutput = io.MultiWriter(os.Stdout, file)
// 	}
// 	log = slog.New(slog.NewJSONHandler(logOutput, &slog.HandlerOptions{AddSource: false})).With(slog.Group("context", "package", "server"))
// }

func server(connection net.Conn, parentChannel chan TCPStatus) {
	// initLog()

	command := make([]byte, 1024)
	for {
		clear(command)
		size, err := connection.Read(command)
		if err != nil {
			if err == io.EOF {
				logger.Log.Info(fmt.Sprintf("Client %s disconnected", connection.RemoteAddr()))
				parentChannel <- TCPStatus{isError: false}
				return
			}
			logger.Log.Error(fmt.Sprint("Failed to read data from connection ", err.Error()))
			parentChannel <- TCPStatus{isError: true}
			return
		}
		logger.Log.Info(fmt.Sprintf("Read %d bytes", size))

		message := respv2.ParseMessage(command[:size])

		logger.Log.Info(fmt.Sprintf("Got following message: %v %T", message, message))

		dispatchV2(message.(respv2.CommandNode), connection)
	}
}

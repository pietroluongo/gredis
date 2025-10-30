package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"

	resp "github.com/codecrafters-io/redis-starter-go/internal/resp"
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
			if err == io.EOF {
				log.Info(fmt.Sprintf("Client %s disconnected", connection.RemoteAddr()))
				parentChannel <- TCPStatus{isError: false}
				return
			}
			log.Error(fmt.Sprint("Failed to read data from connection ", err.Error()))
			parentChannel <- TCPStatus{isError: true}
			return
		}
		log.Info(fmt.Sprintf("Read %d bytes", size))

		message := resp.ParseMessage(command[:size])

		log.Info(fmt.Sprintf("Got following message: %v", message))

		dispatch(message, connection)
	}
}

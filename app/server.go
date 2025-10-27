package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
)

type ValidHandlers string

const (
	Ping ValidHandlers = "ping"
	Pong ValidHandlers = "pong"
	Echo ValidHandlers = "echo"
)

var handlers = map[ValidHandlers]func(net.Conn){
	Ping: handlePing,
	Pong: handlePing,
}

func server(connection net.Conn, parentChannel chan TCPStatus) {
	file, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0644)

	var logOutput io.Writer

	if err != nil {
		slog.Default().Warn(fmt.Sprintf("Failed to create log file, using only stdout %s", err.Error()))
		logOutput = os.Stdout
	} else {
		logOutput = io.MultiWriter(os.Stdout, file)
	}

	log := slog.New(slog.NewTextHandler(logOutput, &slog.HandlerOptions{AddSource: false}))

	command := make([]byte, 1024)
	for {
		size, err := connection.Read(command)
		if err != nil {
			log.Error(fmt.Sprint("Failed to read data from connection ", err.Error()))
			parentChannel <- TCPStatus{isError: true}
			return
		}
		log.Info(fmt.Sprintf("Read %d bytes", size))
		safeCmd := cleanupCommand(command)

		handler := handlers[ValidHandlers(safeCmd)]

		if handler == nil {
			log.Error(fmt.Sprintf("Failed to match handler with cmd '%s' (%v)", safeCmd, []byte(safeCmd)))
			continue
		}
		handler(connection)
		parentChannel <- TCPStatus{isError: false}
		return
	}
}

func handlePing(connection net.Conn) {
	connection.Write([]byte("+PONG\r\n"))
}

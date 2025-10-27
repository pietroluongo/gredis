package main

import (
	"fmt"
	"net"
)

type ValidHandlers string

const (
	Ping ValidHandlers = "PING"
	Pong               = "PONG"
)

var handlers = map[ValidHandlers]func(net.Conn){
	Ping: handlePing,
	Pong: handlePing,
}

func server(connection net.Conn, parentChannel chan TCPStatus) {
	command := make([]byte, 1024)
	for {
		size, err := connection.Read(command)
		if err != nil {
			fmt.Println("Failed to read data from connection ", err.Error())
			parentChannel <- TCPStatus{isError: true}
			return
		}
		fmt.Printf("Read %d bytes\n", size)
		safeCmd := cleanupCommand(command)

		handler := handlers[ValidHandlers(safeCmd)]

		if handler == nil {
			fmt.Printf("Failed to match handler with cmd \"%s\" (%v)\n", safeCmd, []byte(safeCmd))
			connection.Write([]byte("+PONG\r\n"))
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

package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

type TCPStatus struct {
	isError bool
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	c := make(chan TCPStatus)

	slog.Info(fmt.Sprintf("Listening on %s", l.Addr()))

	for {
		connection, err := l.Accept()
		slog.Info(fmt.Sprintf("Got connection from %s", connection.RemoteAddr()))
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go server(connection, c)
	}
}

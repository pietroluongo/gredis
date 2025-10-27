package main

import (
	"fmt"
	"net"
	"os"
)

type TCPStatus struct {
	isError bool
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	c := make(chan TCPStatus)

	for {
		connection, err := l.Accept()
		fmt.Println("got connection")
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go server(connection, c)
		result := <-c
		if result.isError {
			os.Exit(1)
		}
		os.Exit(0)
	}

	// if cleanCommand == "PING" {
	// 	connection.Write([]byte("PONG\r\n"))
	// } else {
	// 	err := fmt.Sprintf("Unrecognized command \"%s\"\r\n", cleanCommand)
	// 	connection.Write([]byte(err))
	// }
}

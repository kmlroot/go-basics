package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Instructions
// go run main.go
// open another terminal and run curl http://localhost:8080
// you should see the result Hello!Hey!
func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Println("listener error")
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		defer conn.Close()

		if err != nil {
			log.Println("connection error")
		}

		io.WriteString(conn, "Hello!")
		fmt.Fprintln(conn, "Hey!")
	}
}

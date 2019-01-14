package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Instructions
// go run main.go
// open another terminal and run telnet localhost 8080
// you should see the result Hello!Hey!
func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Println("listener error")
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("connection error")
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	defer conn.Close()
}

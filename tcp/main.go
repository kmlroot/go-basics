package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Instructions
// go run main.go
// open another terminal and run telnet localhost 8080
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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Println("connection timeout")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	defer conn.Close()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

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
		text := strings.ToLower(scanner.Text())
		bs := []byte(text)
		c := cypher13(bs)

		fmt.Fprintf(conn, "%s - %s \n\n", text, c)
	}

	defer conn.Close()

	fmt.Println("finished")
}

func cypher13(bs []byte) []byte {
	c := make([]byte, len(bs))

	for i, v := range bs {
		if v <= 109 {
			c[i] = v + 13
		} else {
			c[i] = v - 13
		}
	}

	return c
}

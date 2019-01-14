package client

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Println("connection dial error")
	}

	defer conn.Close()

	fmt.Println("connection dialed")
}

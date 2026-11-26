package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func ft_connection(conn net.Conn) {
	defer conn.Close()

	// First message when someone try to connect the server
	conn.Write([]byte("OK hello proto=1\n"))

	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		line := reader.Text()
		switch {
		case line != "CONNECT":
			conn.Write([]byte("COMMANDE INCONNU\n"))
		}
	}

	// Check errors
	err := reader.Err()
	if err != nil {
		fmt.Print("[STDERR]: ", err)
	}
}

func main() {
	// Launch server with 'go run server.go'
	// Join with 'nc localhost 8080'

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go ft_connection(conn)
	}
}

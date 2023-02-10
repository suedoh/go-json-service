package main

import (
	"fmt"
	"net"

    "github.com/suedoh/go-json-service/protocol"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Received connection from %s\n", conn.RemoteAddr().String())

	request := &Request{}
	err := request.FromJSON(conn)
	if err != nil {
		fmt.Printf("Error reading request: %s\n", err)
		return
	}

	fmt.Printf("Received request: %+v\n", request)

	response := NewRequest("response", "hello, client")
	responseJSON, err := response.ToJSON()
	if err != nil {
		fmt.Printf("Error encoding response: %s\n", err)
		return
	}

	_, err = conn.Write(responseJSON)
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error listening: %s\n", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}

		go handleConnection(conn)
	}
}


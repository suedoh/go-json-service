package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/suedoh/go-json-service/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	request := protocol.Request{
		Method: "sum",
		A:      5,
		B:      3,
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(1)
	}

	_, err = conn.Write(requestJSON)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}

	responseReader := bufio.NewReader(conn)
	responseJSON, err := responseReader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	var response protocol.Response
	err = json.Unmarshal(responseJSON, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Result:", strconv.Itoa(response.Result))
}


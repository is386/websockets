package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// This blocks until a connection is made
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to client")

	buffer := make([]byte, 2048)

	for {
		// This blocks until something is received
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		request := strings.Split(string(buffer[:n]), "\r\n")
		requestLine := strings.Split(request[0], " ")
		headers := make(map[string]string)

		for i := 1; i < len(request); i++ {
			header := strings.Split(request[i], ": ")
			if len(header) < 2  {
				continue
			}
			headers[header[0]] = header[1]
		}
		
		if requestLine[0] != "GET" {
			fmt.Println("Error: method of request MUST be GET")
			return
		}
	}
}

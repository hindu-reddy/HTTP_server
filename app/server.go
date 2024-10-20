package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Start listening on the specified port
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	// Accept connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Read the request from the connection
	req := make([]byte, 1024)
	_, err = conn.Read(req)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// Convert request to string
	request := string(req)

	// Extract the request line (first line) by splitting the request string by CRLF
	requestLines := strings.Split(request, "\r\n")
	if len(requestLines) > 0 {
		// Extract the request line components (method, path, and version)
		requestLine := strings.Fields(requestLines[0])

		if len(requestLine) >= 2 {
			// Extract the path
			path := requestLine[1]

			// Check if the path is "/"
			if path == "/" {
				// Respond with 200 OK if the path is "/"
				conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
			} else {
				// Respond with 404 Not Found for any other path
				conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			}
		} else {
			// Malformed request, respond with 404 Not Found
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		}
	}
}

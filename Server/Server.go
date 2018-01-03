package main

import (
	"fmt"
	"net"
	"os"
	StrLib "strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	// Close connection when this function ends
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}
	
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	for{
		fmt.Println("New Connection Recieved: ", conn)
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		s := string(buf[:reqLen])
		fmt.Println("Data Recieved From Client ", s)
		
		if StrLib.Contains(s, "end") == true {
			break
		}
		
		// Send a response back to person contacting us.
		conn.Write([]byte("Message received."))
	}
}
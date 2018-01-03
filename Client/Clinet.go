package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {


	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	fmt.Println("New Connection Recieved: ", conn)
	
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	
	// Send a request to server
	conn.Write([]byte("Data Sent From" + text))

	// Read the incoming message into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	s := string(buf[:reqLen])
	fmt.Println("Data Recieved From Server", s)
	
	// Close the connection when you're done with it.
	conn.Write([]byte("end connection"))
	conn.Close()	  

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
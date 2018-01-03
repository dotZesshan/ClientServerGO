package main

import(
	"net"
	"fmt"
	"bufio"
	"os"
	"time"
	"strconv"
)

func main() {
	for i := 0; i < 50; i++ {
		go MakeClientAndCommunicate(i)
		time.Sleep(time.Second * 5)
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}

func MakeClientAndCommunicate(i int){
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	fmt.Println("New Connection Made with: 127.0.0.1:3333")
	
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	
	// Send a request to server
	t := strconv.Itoa(i)
	conn.Write([]byte("~~Client" + t))

	// Read the incoming message into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	s := string(buf[:reqLen])
	fmt.Println("Data Recieved From Server", s)
	
	// Close the connection when you're done with it.
	conn.Write([]byte("~~Client" + t + "end connection"))
	conn.Close()	  
}
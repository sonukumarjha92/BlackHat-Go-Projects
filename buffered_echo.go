package main

import (
	"bufio"
	"log"
	"net"
)

// echo is an advanced handler using bufio for safe and efficient data buffering.
func echo(conn net.Conn) {
	defer conn.Close()

	// Using bufio to automatically handle buffering instead of manual byte slices
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// Read data until the user presses Enter (newline character '\n')
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected or error occurred.")
			break
		}
		log.Printf("Received data: %s", str)

		// Write the data back to the client
		log.Println("Writing buffered data")
		_, err = writer.WriteString(str)
		if err != nil {
			log.Fatalln("Unable to write data")
		}

		// Flush is mandatory to push the data out of the buffer to the client
		writer.Flush()
	}
}

func main() {
	// Bind to TCP port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	
	log.Println("Listening on 0.0.0.0:20080 (Advanced Buffered Server)")
	
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		
		// Handle connection with goroutine
		go echo(conn)
	}
}


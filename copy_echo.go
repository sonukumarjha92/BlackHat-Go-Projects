package main

import (
	"io"
	"log"
	"net"
)

// echo uses the ultimate 1-liner io.Copy to echo data
func echo(conn net.Conn) {
	defer conn.Close()
	
	// Copy data from io.Reader to io.Writer via io.Copy().
	// It copies from 'conn' and pastes it right back to 'conn'
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	// Bind to TCP port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	
	log.Println("Listening on 0.0.0.0:20080 (Ultimate io.Copy Server)")
	
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		
		go echo(conn)
	}
}


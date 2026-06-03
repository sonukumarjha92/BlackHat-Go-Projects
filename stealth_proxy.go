package main

import (
	"io"
	"log"
	"net"
)

// handle forwards data between the client and the target website
func handle(src net.Conn) {
	// Connect to the target website (Destination)
	dst, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatalln("Unable to connect to target host")
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy source output to destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	
	// Copy destination output back to source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Listen on local port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	
	log.Println("Stealth Proxy Server Listening on 0.0.0.0:20080...")
	
	for {
		conn, err := listener.Accept()
		log.Println("New client connected! Forwarding traffic...")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		
		go handle(conn)
	}
}


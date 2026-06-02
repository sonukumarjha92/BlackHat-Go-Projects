package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	target := "scanme.nmap.org"
	fmt.Println("[*] Initiating SUPER-FAST Scanning on target:", target)
	
	// Record the starting time
	start := time.Now() 

	// WaitGroup to manage concurrent goroutines
	var wg sync.WaitGroup

	// Loop through ports 1 to 1024
	for port := 1; port <= 1024; port++ {
		
		wg.Add(1)

		// Launch a goroutine for each port
		go func(p int) {
			defer wg.Done()
			
			address := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			
			if err != nil {
				// Port is closed or filtered; simply return
				return 
			}
			
			// Port is open
			fmt.Printf("[+] BINGO! Port %d is OPEN\n", p)
			conn.Close()
		}(port)
	}

	// Wait for all goroutines to finish their task
	wg.Wait()
	
	// Print total execution time
	fmt.Printf("[-] Scanning Completed in %v\n", time.Since(start))
}


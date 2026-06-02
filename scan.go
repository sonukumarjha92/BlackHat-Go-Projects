package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// हमारा टारगेट सर्वर और उसका पोर्ट (81 मतलब वेबसाइट का दरवाज़ा)
	target := "scanme.nmap.org:81"
	fmt.Println("Target:", target, "par SYN (Hello) packet bhej rahe hain...")

	// net.DialTimeout बैकग्राउंड में 3-Way Handshake करता है
	// हमने इसे 3 सेकंड का टाइम दिया है (ताकि Filtered Port पर फँस न जाए)
	conn, err := net.DialTimeout("tcp", target, 3*time.Second)

	// अगर Error आया, मतलब दरवाज़ा बंद है (RST) या कोई फायरवॉल है (Timeout)
	if err != nil {
		fmt.Println("❌ Connection Failed! (Port Closed / Filtered)")
		fmt.Println("Error:", err)
		return
	}

	// अगर कोई Error नहीं आया (err == nil), मतलब 3-Way Handshake सफल रहा!
	fmt.Println("✅ SYN-ACK received! Port bilkul OPEN hai! (Handshake Complete 🤝)")
	
	// हैकर मैनर्स: काम होने के बाद कनेक्शन (दरवाज़ा) बंद कर दो
	conn.Close() 
}


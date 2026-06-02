package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "scanme.nmap.org"
	fmt.Println("🚀 Target Scanning Shuru ho rahi hai:", target)

	// हम 1 से लेकर 100 तक के पोर्ट्स स्कैन करेंगे (ताकि रिजल्ट जल्दी मिले)
	for port := 1; port <= 100; port++ {
		
		// IP और Port को जोड़कर एड्रेस बना रहे हैं (जैसे scanme.nmap.org:80)
		address := fmt.Sprintf("%s:%d", target, port)
		
		// सिर्फ 1 सेकंड का टाइमआउट दे रहे हैं ताकि स्कैनिंग तेज़ हो
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err != nil {
			// अगर एरर आया (पोर्ट बंद है), तो हम कुछ नहीं छापेंगे
			// continue का मतलब है - चुपचाप अगले पोर्ट पर चले जाओ
			continue
		}
		
		// अगर एरर नहीं आया, मतलब दरवाज़ा खुला है!
		fmt.Printf("✅ BINGO! Port %d bilkul OPEN hai!\n", port)
		conn.Close()
	}
	
	fmt.Println("🏁 Scanning Complete, Boss!")
}


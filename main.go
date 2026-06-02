package main

import (
	"encoding/json"
	"fmt"
)

// 1. यह हमारे डेटा का ढांचा (Struct) है
type Target struct {
	IP   string
	Port int
}

func main() {
	// 2. हमने Target का डेटा सेट किया
	t := Target{IP: "192.168.1.10", Port: 8080}
	
	// 3. Marshal (पैक करना): डेटा को JSON में बदल रहे हैं
	// (json.Marshal दो चीज़ें देता है: पैक किया हुआ डेटा(b) और कोई एरर(err))
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Packing Error:", err)
		return
	}
	
	fmt.Println("🔒 Network par bheja gaya JSON:")
	fmt.Println(string(b)) // b एक बाइट स्लाइस है, उसे स्ट्रिंग में बदल कर प्रिंट किया

	// 4. Unmarshal (अनपैक करना): JSON को वापस Go डेटा में बदलना
	var newTarget Target
	// यहाँ Pointer (&) का इस्तेमाल किया है
	json.Unmarshal(b, &newTarget) 
	
	fmt.Println("\n🔓 Unpack hone ke baad asli Data:")
	fmt.Println("Target IP:", newTarget.IP)
	fmt.Println("Target Port:", newTarget.Port)
}



package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupMX("instagram.com")
	for _, ip := range iprecords {
		fmt.Println("ip is :", ip)
	}
	// return type is ipv4 and ipv6
}

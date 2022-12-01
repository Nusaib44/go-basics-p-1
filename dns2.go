package main

import (
	"fmt"
	"net"
)

func main() {
	cname, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)
}

//return type is canonical name

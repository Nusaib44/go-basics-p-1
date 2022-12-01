package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go display("world")
	display("hello")
}

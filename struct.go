package main

import "fmt"

type Human struct {
	name  string
	age   int
	place string
}

func main() {
	h1 := Human{"Manu", 22, ""}
	fmt.Println(h1)
	fmt.Println(Human{"Babu", 22, "mm"})
}

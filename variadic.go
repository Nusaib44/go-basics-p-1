package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, " ")
}

func main() {
	var a, b, c string
	fmt.Println("Enter your first name")
	fmt.Scan(&a)
	fmt.Println("Enter your last name")
	fmt.Scan(&b)
	fmt.Println("Enter your initial")
	fmt.Scan(&c)
	d := string("Your name is :")

	// pass a slice in variadic function
	elements := []string{d, a, b, c}
	fmt.Println("")
	fmt.Println(joinstr(elements...))
}

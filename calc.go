package calc

import "fmt"

func Calc() {
	var a, b int

	fmt.Println("Welcome to calculator")
	fmt.Println("Enter value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter value of b: ")
	fmt.Scan(&b)
	ans := a + b
	fmt.Println("Answer = ", ans)
}

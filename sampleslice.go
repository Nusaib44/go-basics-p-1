package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println("cap", cap(a), "   len", len(a))
	fmt.Println("cap", cap(a[:7]), "   len", len(a[:7]))
	//append
	a = append(a, 22, 33, 44) //a=[]int{1,2,3,4,5,6,7,8,22,33,44}
	fmt.Println("After append")
	fmt.Println(a)
	fmt.Println("cap", cap(a), "   len", len(a))
	fmt.Println("cap", cap(a[:7]), "   len", len(a[:7]))
	fmt.Println(a[:])
}

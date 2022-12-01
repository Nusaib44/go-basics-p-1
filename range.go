package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 8, 3, 4, 8, 5, 6, 7, 8, 4}
	for i, e := range a {
		fmt.Println("index :", i, "element", e)
	}
	//check duplicate
	for i, e1 := range a {
		for j, e2 := range a {
			if e2 == -3 {
				continue
			}
			if e1 == e2 && j < i {
				fmt.Println(e2, "  ")
				e2 = -3
			}
		}

	}
}

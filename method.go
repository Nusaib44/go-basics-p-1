package main

import (
	"fmt"
)

type Vertex struct {
	x, y float64
}

func Abc(v Vertex) float64 {

	return (v.x + v.y)
}

func main() {
	var h, j float64
	fmt.Scanln("%f%f", &h, &j)
	v := Vertex{2, 4}
	fmt.Println(Abc(v))
}

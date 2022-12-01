package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade []float64
}

func (m *Student) getage(age int) int {
	m.age = age
	return m.age
}
func (n Student) averageMark() float64 {
	sum := 0.00
	average := 0.00
	for _, m := range n.grade {
		sum += m
		average = sum / float64(len(n.grade))

	}
	return average
}
func main() {
	s1 := Student{"babu", 19, []float64{90, 40, 50}}
	fmt.Println(s1.age)
	s1.getage(14)
	fmt.Println(s1.age)
	fmt.Println(s1.averageMark())
}

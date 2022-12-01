package main

import (
	"fmt"
	"time"
)

func main() {
	//Current Time
	fmt.Println(time.Now())
	//Time
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	//Date
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	//condition
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println("Today is :", today)
}

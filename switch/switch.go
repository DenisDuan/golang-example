package main

import (
	"fmt"
	"time"
)

func main() {
	switchInteger()
	switchWeek()
	switchTime()
	switchType(true)
	switchType(12.2)
	switchType("Hello")
	switchType(12)
}

func switchInteger() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func switchWeek() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("It's weekday :-(")
	}
}

func switchTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's morning now.")
	default:
		fmt.Println("It's afternoon")
	}
}

func switchType(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("This is a boolean value")
	case int:
		fmt.Println("This is an integer value")
	case string:
		fmt.Println("This is a string value")
	default:
		fmt.Println("Don't know what type it is: ", t)
	}
}

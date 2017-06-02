package main

import "fmt"

func main() {
	// variable can be initialize with value
	if num := 9; num < 0 {
		fmt.Println(num, " is a negative number")
	} else if num < 10 {
		fmt.Println(num, " only has one digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}

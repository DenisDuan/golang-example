package main

import "fmt"

// This is a demonstration of using recursive call to perform factorial a number

func main() {
	fmt.Println("The factorail of 7 is: ", factorial(7))
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
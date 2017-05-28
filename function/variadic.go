package main

import "fmt"

func main() {
	fmt.Println("The plusThreeNums of 1, 2, 3, 4, 5, 6 is: ", sum(1, 2, 3, 4, 5, 6))
}

func sum(nums ...int) int {

	total := 0

	for _, arg := range nums {
		total = total + arg
	}

	return total
}

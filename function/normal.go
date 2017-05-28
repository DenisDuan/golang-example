package main

import "fmt"

func main() {
	fmt.Println("value return by plus(1, 2);", plus(1, 2))
	fmt.Println("value return by plusThreeNums(1, 2, 3);", plusThreeNums(1, 2, 3))
}

func plus(a, b int) int {
	return a + b
}

func plusThreeNums(a, b, c int) int {
	return a + b + a
}

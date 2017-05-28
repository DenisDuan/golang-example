package main

import "fmt"

func main() {
	var a, b = getVals(1, 2)
	fmt.Println("The values returned of given 1, 2 are: ", a, b)

	var _, d = getVals(3, 4)
	fmt.Println("The vlues returned of given 3, 4 are:", d)
}

func getVals(a, b int) (int, int) {
	return a + 1, b + 2
}

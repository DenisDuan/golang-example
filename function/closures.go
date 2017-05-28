package main

import "fmt"

func main() {
	firtAddOneFunc := addOne()

	fmt.Println(firtAddOneFunc())
	fmt.Println(firtAddOneFunc())
	fmt.Println(firtAddOneFunc())

	// The second closure is closing another function
	secondAddOneFunc := addOne()
	fmt.Println(secondAddOneFunc())
	fmt.Println(secondAddOneFunc())
	fmt.Println(secondAddOneFunc())

}

func addOne() func() int {
	i := 0

	return func() int {
		i = i + 1
		return i
	}
}

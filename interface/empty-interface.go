package main

import "fmt"

/*
An empty interface may hold values of any type (Every type implements at least zero methods)
 */
func main() {
	var i interface{}
	describeInterface(i)

	i = 42
	describeInterface(i)

	i = "hello"
	describeInterface(i)
}

func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

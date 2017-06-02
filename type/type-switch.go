package main

import (
	"fmt"
)

/*
The declaration in a type switch has the same syntax as a type assertion i.(T),
but the specific type T is replaced with the keyword "type".
*/

func doTypeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	doTypeSwitch(21)
	doTypeSwitch("Hello")
	doTypeSwitch(true)
}

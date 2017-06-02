package main

/*
A type assertion provides access to an interface value's underlying concrete value
t := i.(T)
*/
import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//f, ok := i.(float64)
	//fmt.Printf(f, ok)

	//f = i.(float64)
	//fmt.Printf(f)

}

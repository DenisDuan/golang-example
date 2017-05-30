package main

import "fmt"

func main() {

	// Create struct with default
	fmt.Println(person{"Denis", 18})

	// Create struct using specific fields
	fmt.Println(person{name: "Denis Two", age: 16})

	// Create struct with missing filed. The missing field will be default value
	fmt.Println(person{name: "missing age"})
	fmt.Println(person{age: 16})

	// add & in front of struct gets the pointer to the struct
	fmt.Println(&person{"Tom Cruise", 66})

	// access the field value of a pointer of a struct will be dereferenced automatically
	p := person{"Peter", 56}
	sp := &p
	fmt.Println(sp.name)

	// structs are mutable
	p.age = 18
	fmt.Println(p)

}

type person struct {
	name string
	age  int
}

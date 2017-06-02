package main

import (
	"fmt"
)

func f(from string) {
	for i:=0;i<3;i++ {
		fmt.Println(from, ":", i)
	}
}

func main()  {
	f("direct")

	go f("go routine")

	// Anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("go anonymously")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Finished")
}


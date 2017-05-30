package main

import "fmt"

func main() {

	i := 1

	fmt.Println("The original value of i is: ", i)

	zeroVal(i)
	fmt.Println("After passing into zeroVal(), i is: ", i)

	zeroPtr(&i)

	fmt.Println("After passing into zeroPtr(), i is: ", i)

	fmt.Println("The memory address of i is: ", &i)

}

/*
Function is passed by value, zeroVal will get a copy of intVal
distinct from the one in the calling function.
*/
func zeroVal(intVal int) {
	intVal = 0
}

/*
zeroPtr in contrast has an *int parameter, meaning that it takes an int pointer.
The *intPtr code in the function body then dereferences the pointer from its memory
address to the current value at that address. Assigning a value to a dereferenced
pointer changes the value at the referenced address.
*/
func zeroPtr(intPtr *int) {
	*intPtr = 0
}

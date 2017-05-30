package main

import "fmt"

func main() {
	r := retangular{10, 20}
	fmt.Println("The area of retangular using value reciever: ", r.area())
	fmt.Println("The perimeter of retangular using value reciever: ", r.perimeter())

	rp := &r
	fmt.Println("The area of retangular using point reciever: ", rp.area())
	fmt.Println("The perimeter of retangular using point reciever: ", rp.perimeter())
}

type retangular struct {
	width, height int
}

/*
It's better to use pointer receiver type to avoid copying on method calls
or to allow the method to mutate the receiving struct.
*/
func (r *retangular) area() int {
	return r.width * r.height
}

func (r retangular) perimeter() int {
	return 2*r.width + 2*r.height
}

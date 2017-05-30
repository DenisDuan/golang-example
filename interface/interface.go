package main

import (
	"fmt"
	"math"
)

func main() {
	rec := rectangular{"rectangular", 10, 20}
	cir := circle{"circle", 10}
	measure(rec)
	measure(cir)
}

type geometry interface {
	getName() string
	area() float64
	perimeter() float64
}

// The struct and methods  of rectangular
type rectangular struct {
	name          string
	width, height float64
}

func (r rectangular) getName() string {
	return r.name
}

func (r rectangular) area() float64 {
	return r.width * r.height
}

func (r rectangular) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// The struct and method of circle
type circle struct {
	name   string
	radius float64
}

func (c circle) getName() string {
	return c.name
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g.getName())
	fmt.Println("The area is: ", g.area())
	fmt.Println("The perimeter is: ", g.perimeter())
}

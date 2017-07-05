package main

import (
	"fmt"
	"math"
)

// define our type Square which takes the length of a side
type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

// attach area method to our type Square
func (z Square) area() float64 {
	return z.side * z.side
}

// attach area method to our type Circle
// now this can also be used as type Shape since it has the area method!
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// any type that implements the area() float64 method will will be a type Shape
type Shape interface {
	area() float64
}

func info(z Shape) { // this only takes in a type Shape object!
	fmt.Println(z)
	fmt.Println(z.area())
}
func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}

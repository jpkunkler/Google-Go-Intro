package main

import "fmt"

// Square defines our type Square which takes the length of a side
type Square struct {
	side float64
}

// Shape describes that any type that uses the area() float64 method will be a shape
type Shape interface {
	area() float64
}

// attach method to our type Square
func (z Square) area() float64 {
	return z.side * z.side
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}
func main() {
	s := Square{10}
	info(s)
}

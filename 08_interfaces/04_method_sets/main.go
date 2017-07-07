package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// if we have a pointer receiver here, we have to pass in an address
func (s *square) area() float64 {
	return s.side * s.side
}

// if we only take in a shape, we can pass either an address or a value!
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{5}
	c := circle{5}
	info(&s)
	info(c)
	info(&c)
}

/*

RECEIVERS	VALUES
-------------------------
(t T)		T and *T
(t *T)		*T

VALUES		RECEIVERS
--------------------------
T		(t T)
*T		(t T) AND (t *T)  // A POINTER VALUE WORKS WITH BOTH!!

*/

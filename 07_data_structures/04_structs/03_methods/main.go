package main

import "fmt"

// we declare our type here
type person struct {
	first string
	last string
	age int
}

// we create our function with a receiver (p persons) which connects this func
// to our type and makes it a method of this type
func (p person) fullName() string {
	return p.first + " " + p.last
}

func (p person) doubleAge() int {
	return p.age * 2
}

func main() {
	p1 := person{"Jan", "Kunkler", 22}
	fmt.Println(p1.fullName())
	fmt.Println(p1.doubleAge())
}

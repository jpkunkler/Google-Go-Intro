package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

// DoubleZero is the outer type
type DoubleZero struct {
	Person // Person is the inner type
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

// if the receiver is a double zero, this will override the regular person method
func (dz DoubleZero) Greeting() {
	fmt.Println("Miss MoneyPenny, it's so good to see you.")
}
func main() {
	p1 := DoubleZero{
		Person{
			"James",
			"Bond",
			32,
		},

		true,
	}

	// just a regular person
	p2 := Person{"Jan", "Kunkler", 22}

	p1.Greeting()
	p2.Greeting()
	p1.Person.Greeting() // even though it's James Bond, we refer to just the regular Person (inner type)
}

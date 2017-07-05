package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

// DoubleZero is the outer type
type DoubleZero struct {
	Person // Person is the inner type
	first string // this will override person's first field because outer > inner
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person{
			"James",
			"Bond",
			32,
		},

		"Double Zero Seven", // here we override our previous name
		true,
	}
	// Double Zero Seven will be our first name instead of James
	fmt.Println(p1.first, p1.last, p1.age, p1.LicenseToKill)
}
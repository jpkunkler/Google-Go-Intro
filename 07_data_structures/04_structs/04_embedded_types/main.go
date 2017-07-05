package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
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
	fmt.Println(p1.first, p1.last, p1.age, p1.LicenseToKill)
}
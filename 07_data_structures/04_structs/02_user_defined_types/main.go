package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 22
	fmt.Println(myAge)

	var yourAge int
	yourAge = 29
	fmt.Println(yourAge)

	// this will not work
	// invalid operation: myAge + yourAge (mismatched types foo and int)
	// fmt.Println(myAge + yourAge)

	// this will work
	fmt.Println(int(myAge) + yourAge)
}

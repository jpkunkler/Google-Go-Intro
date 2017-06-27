package main

import "fmt"

func wrapper() func() int { // the return of wrapper() will be type "func() int" -- THAT'S A TYPE! An actual function will be returned.
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper() //increment will now be a function and can be used as such.
	fmt.Println(increment())
	fmt.Println(increment())
}

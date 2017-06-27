package main

import "fmt"

func main() {

	// func expressions are the only way to define functions within functions!
	// a function is basically assigned to a variable
	greeting := func() {
		fmt.Println("Hello World!")
	}

	greeting()
	// type of greeting is func()
}

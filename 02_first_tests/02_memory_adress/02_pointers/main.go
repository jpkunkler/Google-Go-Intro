package main

import "fmt"

func main() {

	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b int = &a // b is of type "int pointer" *int -- the * is part of the type!!

	fmt.Println(b)
	fmt.Println(*b) // *b means de-referencing memory address

	*b = 44 // the value at this address, change it to 44
	fmt.Println(a)
}

package main

import "fmt"

func withPointer(x *int) {
	*x = 0
}

func noPointer(x int) {
	x = 0
}

func main() {
	x := 5
	noPointer(x)
	fmt.Println("Using no Pointer:")
	fmt.Println(x)
	withPointer(&x)
	fmt.Println("Using a Pointer:")
	fmt.Println(x)

}

/*
Assertion only works with interfaces
since conversion cannot be used with interfaces!
*/

package main

import "fmt"

func main() {
	var x interface{} = 6 // x is of type empty interface

	fmt.Printf("%T - %d \n", x, x)
	//x = float64(x) // this conversion will not work
	//x = x.(float64) // this does also not work

	// this is an assertion!
	// this will work but set f to the zero value
	// since x is not of type float64 but int
	f, ok := x.(float64)
	fmt.Printf("%T - %f - %v \n", f, f, ok)

	t, ok := x.(int)
	fmt.Printf("%T - %d - %v \n", t, t, ok)
}

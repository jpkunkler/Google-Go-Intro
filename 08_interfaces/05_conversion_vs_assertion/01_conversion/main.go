/*
Conversion is basically just converting from one type to another
for example from int to float64 --> widening
from float64 to int --> narrowing (leaving out the decimal places)
from rune to string and vice versa
*/

package main

import "fmt"

func main() {
	var x int = 6
	var y float64 = 32.156
	fmt.Printf("x: %T - %d \n", x, x)
	fmt.Printf("y: %T - %f \n", y, y)
	fmt.Println("Converting...")
	fmt.Printf("x: %T - %f \n", float64(x), float64(x))
	fmt.Printf("y: %T - %d \n", int(y), int(y))
}

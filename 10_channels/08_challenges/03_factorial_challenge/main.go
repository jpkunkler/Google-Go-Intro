/*
Rewrite the program below using goroutines for concurrency
*/
package main

import "fmt"

func main() {
	c := factorial(4)
	fmt.Println(c)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i

	}
	return total
}

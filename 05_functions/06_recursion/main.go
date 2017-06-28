package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // factorial is called by itself --> recursion! A Function calling itself
}

func main() {
	fmt.Println(factorial(5))
}

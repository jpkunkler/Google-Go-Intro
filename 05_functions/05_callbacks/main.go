package main

import "fmt"

// #FUNCTIONAL PROGRAMMING
func visit(numbers []int, callback func(int)) { // callback is used as a PARAMETER!!
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 5, 8, 13}, func(n int) { // for callback to work, a function needs to be passed as second argument.
		fmt.Println(n)
	})
}

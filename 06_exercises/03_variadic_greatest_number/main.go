package main

import "fmt"

func greatest(numbers ...int) int {
	var x int
	for _, n := range numbers {
		if n > x {
			x = n
		}
	}
	return x
}

func main() {
	result := greatest(1, 2, 3, 4, 2, 10, 231, 7)
	fmt.Println(result)
}

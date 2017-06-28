package main

import "fmt"

// 2 PARAMETERS:
// 1.) Slice of int (list of integers)
// 2.) A Function that takes in an integer and returns a boolean
// Finally, a Slice of int will be returned.
func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{} // create empty list
	for _, n := range numbers {
		if callback(n) { // if callback function returns true
			xs = append(xs, n) // append number to list xs
		}
	}
	return xs
}

func main() {
	result := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(result)
}

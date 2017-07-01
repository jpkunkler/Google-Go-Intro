package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4]) // slicing a slice!
	fmt.Println(mySlice[4])   // index a slice
}

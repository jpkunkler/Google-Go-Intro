package main

import "fmt"

func main() {

	mySlice := make([]int, 1)
	fmt.Println(mySlice)
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++ // increment like this!
	fmt.Println(mySlice[0])
	mySlice[0] += 100 // add x amount to mySlice
	fmt.Println(mySlice[0])
}

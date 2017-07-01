package main

import "fmt"

// once the capacity is full, it will be doubled automatically!
func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", i)
	}
}

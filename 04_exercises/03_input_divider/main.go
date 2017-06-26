package main

import "fmt"

func main() {
	var small int
	var large int

	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)

	fmt.Print("Please enter a large number: ")
	fmt.Scan(&large)

	result := large % small

	fmt.Println("The remainder is:", result)
}

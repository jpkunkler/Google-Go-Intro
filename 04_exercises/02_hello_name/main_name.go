package main

import "fmt"

func main() {
	var name string
	fmt.Print("Who Am I: ")
	fmt.Scan(&name)

	fmt.Println("Hello, my name is", name)
}

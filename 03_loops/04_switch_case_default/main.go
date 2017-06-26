package main

import "fmt"

func main() {

	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	switch name {
	case "Daniel":
		fmt.Println("Wassup Daniel?!")
	case "Jan":
		fmt.Println("Wassup Jan?!")
	case "Aden":
		fmt.Println("Wassup Aden?!")
	default:
		fmt.Println("I guess I don't know your name..")
	}
}

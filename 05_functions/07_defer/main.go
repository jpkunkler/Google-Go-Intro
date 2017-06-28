package main

import "fmt"

// DEFER is really useful in combination with opening and closing files. You simply defer the closing call and it will be run right at the end, but your open and close code will be right in one block.
func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // defer will delay the action until the end of the function. This will run right before func main() finishes!
	hello()
}

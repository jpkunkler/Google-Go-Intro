package main

import "fmt"

func main() {
	var x [20]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[2] = 3
	fmt.Println(x)
}

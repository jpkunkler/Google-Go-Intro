package main

import "fmt"

const (
	_  = iota             // throw away 0
	kb = 1 << (iota * 10) // bitshift to the left: 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2*10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Printf("%b ", kb)
	fmt.Printf("%d\t\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\t\n", mb)

}

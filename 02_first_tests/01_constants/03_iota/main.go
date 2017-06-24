package main

import "fmt"

const (
	_  = iota // throw away 0
	KB = 1 << (iota * 10) // bitshift to the left: 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Printf("%b ", KB)
	fmt.Printf("%d\t\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t\n", MB)
	
}

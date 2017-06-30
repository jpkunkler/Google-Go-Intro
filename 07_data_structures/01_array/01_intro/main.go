package main

import "fmt"

// an array is not dynamic. It has a fixed length and will always keep this length.
func main() {
	var a [58]string // an array has a fixed length!! in this case 58
	fmt.Println(a)
	fmt.Println(a[48])
	for i := 65; i <= 122; i++ {
		a[i-65] = string(i)
	}
	fmt.Println(a)
	fmt.Println(a[48])
}

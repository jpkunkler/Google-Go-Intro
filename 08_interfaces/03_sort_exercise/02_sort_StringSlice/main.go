package main

import (
	"fmt"
	"sort"
)
func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"} // this is not a user-defined type!

	fmt.Println(s)
	sort.StringSlice(s).Sort() // StringSlice attaches the methods of Interface to []string, sorting in increasing order.
	fmt.Println(s)
}

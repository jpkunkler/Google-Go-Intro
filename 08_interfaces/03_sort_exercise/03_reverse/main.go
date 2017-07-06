package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1,3,29,7,32,16}

	sort.Sort(sort.Reverse(sort.IntSlice(s))) // IntSlice attaches the methods of Interface to []int, sorting in increasing order.
	fmt.Println(s)
}

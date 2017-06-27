package main

import "fmt"

func main() {
	data := []float64{2, 6} // this is a slice of float64 (basically a list/array)
	n := average(data...)   // ... behind indicates a variadic argument, meaning that every item (in that slice) will be past by itself
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

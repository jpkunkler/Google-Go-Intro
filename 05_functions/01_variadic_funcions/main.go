package main

import "fmt"

func main() {
	fmt.Println(average(2, 6))               // 2 input arguments
	fmt.Println(average(15, 29, 30, 87, 52)) // 5 input arguments
}

func average(sf ...float64) float64 { // ... (three dots) indicates variadic input. Meaning this parameter can take any number of input arguments.
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

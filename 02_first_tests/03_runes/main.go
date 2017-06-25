package main

import "fmt"

func main() {

	for i := 0; i <= 1000; i++ {

		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}

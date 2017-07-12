package main

import (
	"fmt"
)

func main() {
	c := gen(2, 3, 16, 32)
	out := sq(c)

	for n := range out { // this takes off from the result channel 'out'
		fmt.Println(n)
	}
}
func gen(nums ...int) chan int { // this will take a variadic amount of integers and puts them onto channel
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out // then it returns the channel
}

func sq(input chan int) chan int { // this takes the above channel as input
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n // does our calculations and puts them onto a channel
		}
		close(out)
	}()
	return out // and finally returns our channel containing results
}

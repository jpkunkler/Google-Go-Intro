package main

import "fmt"

func main() {
	c := factorial(4)  // factorial will return a channel of type int
	for n := range c { // loop over channel c
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int) // create our channel
	go func() {           // launch our goroutine in the background
		total := 1
		for i := n; i > 0; i-- {
			total *= i

		}
		out <- total // put our result onto the channel
		close(out)   // close the channel once we are done
	}()
	return out // return a channel
}

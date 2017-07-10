package main

import "fmt"

func main() {
	c := make(chan int) // create an unbuffered channel for integers

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // pass a value to the channel
		}
		close(c) // close the channel so nothing can be put into it --> contained values can still be received!
	}()

	for n := range c { // print until channel is closed
		fmt.Println(n)
	}
}


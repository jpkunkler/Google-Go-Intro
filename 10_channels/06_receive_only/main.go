package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int { // this is now a read/receive-only channel!
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int { // puller also only takes a receive-only chan and returns the same
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
The <- in front of a channel makes it a RECEIVE-ONLY channel
No values an be put onto the channel, but all values can be received from it
*/

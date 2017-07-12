package main

import "fmt"
import "sync"

func main() {
	in := gen(2, 3)

	// FAN OUT
	// distribute the sq work across two goroutines that both read from them
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

// Generate and populate a channel with numbers
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// Square input values and return a channel
func sq(input chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int { // this takes in a []chan int --> multiple channels!
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// use a func expression for readability
	// fetch output from channel and send it to single output channel
	output := func(ch chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	// loop over every channel in our []chan int
	for _, c := range cs {
		go output(c)
	}

	// Wait until every goroutine is done, then close the channel to prevent deadlock
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

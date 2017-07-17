package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(1000) // create 1.000 batches á 10 numbers --> 10.000 values

	// FAN OUT 10.000 calcuations to 10 workers
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	var count int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		count++
		fmt.Println(count, "\t", n)
	}
}

// gen() returns a receive-only channel populated with integers
func gen(batches int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < batches; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

// factorial() creates the workerProcess for a factorial calculation
func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

// fact() returns the factorial value of an integer
func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// FAN IN
// merge() collects values from a variadic amount of channels and puts them onto 1 output channel
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// take values off a channel and put them onto the output channel
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// run goroutines for all input channels
	for _, n := range cs {
		go output(n)
	}

	// wait until all our goroutines are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern
CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/

package main

import "fmt"

func main() {
	values := gen()        // values will be a receive-only channel
	c := factorial(values) // pass our receive-only channel and retrieve the new one as c
	for n := range c {     // take all results off from channel c
		fmt.Println(n)
	}
}

func gen() <-chan int { // this generates 100 values and puts them onto a receive-only channel
	out := make(chan int) // create our out channel
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out // return our channel
}

func factorial(input <-chan int) <-chan int { // also returnsa receive-only channel
	out := make(chan int)
	go func() {
		for n := range input {
			out <- fact(n) // call function fact for every value in input channel
		}
		close(out)
	}()
	return out
}

func fact(n int) int { // this is our factorial computation: 4! = 4 * 3 * 2 * 1
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

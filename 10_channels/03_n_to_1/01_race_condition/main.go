package main

import (
	"sync"
	"fmt"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1) // this causes a race condition because both functions access it!
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1) // this causes a race condition because both functions access it!
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait() // wait for goroutines to finish
		close(c) // close channel after all is done
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
	}()
}

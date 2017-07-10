package main

import (
	"sync"
	"fmt"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) // FIX: add to WaitGroup outside of goroutines to prevent race condition!
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
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

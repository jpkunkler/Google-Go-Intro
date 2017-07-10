package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true // set channel done to true if finished
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done // check if first is done (if true)
		<-done // check if second is done (if true)
		close(c) // close channel after all is done
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
	}()
}

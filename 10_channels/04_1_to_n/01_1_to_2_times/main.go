package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // create channel to store values
	done := make(chan bool) // create channel to store status

	go func ()  {
		for i := 0; i < 1000; i++ {
			c <- i // put value onto channel
		}
		close(c)
	}()

	go func ()  { // this func will take values off the channel
		for n := range c {
			fmt.Println(n)
		}
		done <- true // set channel to true if finished
	}()

	go func ()  { // this will also take values off the channel
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	// wait for both goroutines to finish
	<-done
	<-done
}
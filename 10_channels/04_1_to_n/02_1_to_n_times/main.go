package main

import (
	"fmt"
)

func main() {
	const n = 10 // define how many goroutines/functions we want
	c := make(chan int)
	done := make(chan bool)

	go func ()  { // same as before, populate the channel
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ { // create n amount of goroutines
		go func ()  {
			for n := range c { // every goroutine accesses channel c
				fmt.Println(n)
			}
			done <- true // every goroutine has to communicate that it's done
		}()
	}

	for i := 0; i < n; i++ { // wait for all n functions to return done
		<-done
	}
}
/*
CHANNELS
work like relay racers
One has to pass on the baton for the other to continue
they have to synchronize

the first function will only continue once the second function has printed the value
and thus taken it off from the channel
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // create a channel for integers

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // pass a value to the channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // once this has printed, above func will continue
		}
	}()

	time.Sleep(time.Second)
}

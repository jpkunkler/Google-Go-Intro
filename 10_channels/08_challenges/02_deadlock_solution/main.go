/*
This program is supposed to print the numbers 0 to 10 but instead only returns 0.
Why is that? Identify the deadlock and fix it.
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // we need to close our channel once we're done, otherwise there's a deadlock
	}()

	for n := range c { // we need to loop over our channel to retrieve all values from it
		fmt.Println(n)
	}
}

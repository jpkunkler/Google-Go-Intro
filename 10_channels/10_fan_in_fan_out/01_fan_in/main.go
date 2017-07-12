package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ { // this will limit the channel to only take the first 10 values
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring; I'm leaving.")

	// ###########################
	// IMPORTANT:
	// this part below would end in an infinite loop program! We are not calling close() anywhere
	// so our channel c would just keep on giving us new values
	// which is why we need to limit the amount we are pulling off from it
	// like we are doing above!
	// ###########################
	// for n := range c {
	// 	fmt.Println(n)
	// }
}

func boring(msg string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

// FAN IN (take values off from 2 channels and continue as one)
func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1 // take value off from input1 and onto channel out
		}
	}()
	go func() {
		for {
			out <- <-input2
		}
	}()
	return out
}

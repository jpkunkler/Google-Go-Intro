/*
DATA RACE
if concurrent processes try to access the same value, there will be data race
meaning that they override each other

use

go run -race main.go

to check for race conditions
*/
package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// assign and instantiate a WaitGroup
var wg sync.WaitGroup

// both functions will try to assign values to test concurrently --> race condition
var counter int

func main() {
	wg.Add(2)

	go incrementor("Foo: ")
	go incrementor("Bar: ")

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond) // used to artificially give time for other process
		counter = x
		fmt.Println(s, i, "Counter: ", counter)

	}
	wg.Done()
}
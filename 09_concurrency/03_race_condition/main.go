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
)

// assign and instantiate a WaitGroup
var wg sync.WaitGroup

// both functions will try to assign values to test concurrently --> race condition
var test int

func main() {
	wg.Add(2) // add 2 to WaitGroups since we run 2 functions

	// start goroutines
	go foo()
	go bar()
	wg.Wait() // wait until all functions are done
}

func foo() {
	for i := 0; i < 45; i++ {
		test = i
		fmt.Println("Foo:", test)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		test = i
		fmt.Println("Bar:", test)
	}
	wg.Done()
}

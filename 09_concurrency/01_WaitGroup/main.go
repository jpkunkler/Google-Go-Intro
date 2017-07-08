/*
CONCURRENCY
doing many things, but only one at a time
--> "multitasking"
*/
package main

import (
	"fmt"
	"sync"
)

// assign and instantiate a WaitGroup
var wg sync.WaitGroup

func main() {
	wg.Add(2) // add 2 to WaitGroups since we run 2 functions

	// start goroutines
	go foo()
	go bar()
	wg.Wait() // wait until all functions are done
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}

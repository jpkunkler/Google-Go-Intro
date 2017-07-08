/*
ATOMICITY
 */
package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64 // int64 specifically needed for atomic (or int32)


func main() {
	wg.Add(2)

	go incrementor("Foo: ")
	go incrementor("Bar: ")

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// this will prevent race conditions
		atomic.AddInt64(&counter, 1) // we need to pass in the address to a var
		fmt.Println(s, i, "Counter: ", atomic.LoadInt64(&counter)) // access without race

		// CAUTION!!!
		// fmt.Println(s, i, "Counter: ", counter) >>> THIS WOULD CAUSE A DATA RACE!!
	}
	wg.Done()
}
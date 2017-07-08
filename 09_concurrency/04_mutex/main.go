/*
MUTEX stands for mutual exclusion object
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)

	go incrementor("Foo: ")
	go incrementor("Bar: ")

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// lock the process so that no other thread can use it
		mutex.Lock()

		// do the work here while only we have access
		counter++
		fmt.Println(s, i, "Counter: ", counter)

		// unlock it so others get to access it again
		mutex.Unlock()
	}
	wg.Done()
}

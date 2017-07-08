/*
This would be an example of no concurrency
There are no goroutines used
Those 2 functions will be run one after the other

Foo will run first until it's done
Bar will run as soon as Foo is finished.
 */

package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
	}
}
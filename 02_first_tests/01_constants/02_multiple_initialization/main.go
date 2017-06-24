package main

import "fmt"

const (
	pi       = 3.14
	language = "Go"
	name string = "Jan"
)

func main() {
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Printf("%T\t%q",name, name)
}

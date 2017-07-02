package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 10

	fmt.Println(m)
	delete(m, "k2")
	fmt.Println(m)

	_, ok := m["k2"]
	fmt.Println("ok?", ok)

	// direct initialization
	m2 := map[string]int{"test1": 42, "test2": 4923}
	fmt.Println(m2)
}

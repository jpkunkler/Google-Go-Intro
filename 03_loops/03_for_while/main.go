package main

import "fmt"

func main() {

	i := 0
	for i < 10 { // there is no "while" in Go - just use for!
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		if i >= 20 {
			break
		}
		i++
	}
	for {
		i++
		if i%2 == 0 { // check if i is even. If so, then continue --> ONLY ODD NUMBERS WILL BE DISPLAYED!
			continue
		}

		fmt.Println(i)

		if i >= 50 {
			break
		}
	}
}

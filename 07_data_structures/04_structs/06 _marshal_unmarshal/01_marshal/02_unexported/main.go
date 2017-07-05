package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// those fields are not exported (lowercase)
	first string
	last string
	age int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(p1)
	fmt.Println(string(bs)) // this will be empty since nothing is exported from our type
}

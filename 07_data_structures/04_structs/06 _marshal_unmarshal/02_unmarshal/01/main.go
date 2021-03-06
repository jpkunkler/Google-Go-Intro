package main

import (
	"encoding/json"
	"fmt"
)

// Person is an entry for a single person
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	// this person is initialized with zero values
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First": "James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1) // unmarshal those values to memory address of struct p1

	fmt.Println("---------------- AFTER UNMARSHAL:")
	// now p1 has the extracted values!
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}

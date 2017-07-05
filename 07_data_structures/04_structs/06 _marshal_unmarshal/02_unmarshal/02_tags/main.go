package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string `json:"-"`// we will not have this field in our input values
	Age int `json:"wisdom-score"`
}

func main() {
	// this person is initialized with zero values
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)


	bs := []byte(`{"First": "James", "wisdom-score":20}`) // no Last Name!
	json.Unmarshal(bs, &p1) // unmarshal those values to memory address of struct p1

	fmt.Println("---------------- AFTER UNMARSHAL:")
	// now p1 has the extracted values!
	fmt.Println(p1.First)
	fmt.Println(p1.Last) // last will still be an empty string (default for strings)
	fmt.Println(p1.Age)

}
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// those fields are not exported (lowercase)
	First string
	Last string `json:"-"` // this tag will exclude this value for our json export
	Age int `json:"wisdom-score"` // this value will be renamed in our json export
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(p1)
	fmt.Println(string(bs))
}

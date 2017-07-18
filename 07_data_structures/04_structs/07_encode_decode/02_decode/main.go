package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person is an entry for a single person
type Person struct {
	First       string
	Last        string
	Age         int `json:"wisdom-score"`
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First": "James", "Last":"Bond","wisdom-score":20}`) // simulate json stream coming in
	json.NewDecoder(rdr).Decode(&p1)                                                // decode into memory address of struct p1

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}

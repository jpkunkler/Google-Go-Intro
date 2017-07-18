package main

import (
	"encoding/json"
	"os"
)

// Person is an entry for a single person
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Jan", "Kunkler", 22, 183}
	json.NewEncoder(os.Stdout).Encode(p1) // encode our struct to json and print to terminal
}

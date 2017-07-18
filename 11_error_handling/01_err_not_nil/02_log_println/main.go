package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}

/*
The log.Println acts in the same way as its fmt counterpart but ADDS A DATETIME to it.
*/

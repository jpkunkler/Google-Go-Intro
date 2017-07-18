package main

import (
	"fmt"
	"log"
	"os"
)

// this init function is used to set up our log file
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // declares which file to use as log file
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err) // if a logfile is setup, this will write to StdErr instead of StdOut
	}
}

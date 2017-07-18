package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err) // this will cause the programm to panic and quit with exit status 2
	}
}

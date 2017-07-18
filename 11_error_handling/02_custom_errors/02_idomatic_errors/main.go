package main

import (
	"errors"
	"log"
)

// all errors are declared at the same point
// all errors begin with Err...
var (
	ErrNegativeSquare = errors.New("square root of negative number")
	ErrNorgateMath    = errors.New("Impossible math to do")
)

func main() {
	_, err := sqrt(-17)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSquare // simple reference the error variable here
	}

	return 42, nil
}

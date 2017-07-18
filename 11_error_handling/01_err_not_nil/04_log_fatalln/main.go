package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("err happened", err)
	}
}

/* log.Fatalln does the same output as fmt.Println but also returns exit status 1 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close() //after we are done, close our response

	words := make(map[string]string)

	// scan our result into buffer
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords) // split buffer by words

	for sc.Scan() { // for every word we scanned
		words[sc.Text()] = "" //add entry to our map
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words { //we only need the key since our values are all empty ""
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}

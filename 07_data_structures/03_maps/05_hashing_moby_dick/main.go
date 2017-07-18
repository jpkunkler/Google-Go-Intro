package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// retrieve moby dick text file
	res, err := http.Get("http://algs4.cs.princeton.edu/63suffix/mobydick.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// load result into buffer
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for our scanning operation
	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12) // we have 12 buckets containing slices of strings i.e. our words

	// Create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	// Loop over all words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)             // hash every word using our algorithm
		buckets[n] = append(buckets[n], word) // add our word to the appropriate bucket
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// Print words in one of the buckets
	//fmt.Println(buckets[6])
}

func hashBucket(word string, buckets int) int {
	var sum int
	// take the sum of letters for each word
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets // this will create a more evenly distributed set of buckets
}

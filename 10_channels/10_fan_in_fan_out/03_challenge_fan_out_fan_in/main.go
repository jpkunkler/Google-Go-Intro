package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var workerID int64
var publisherID int64

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(5 * time.Millisecond)
}

func publisher(out chan string) {
	atomic.AddInt64(&publisherID, 1)
	thisID := atomic.LoadInt64(&publisherID)
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	atomic.AddInt64(&workerID, 1)
	thisID := atomic.LoadInt64(&workerID)
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

/*
CHALLENGE #1:
Is this fan out?
YES
*/

/*
CHALLENGE #2:
Is this fan in?
NO
*/

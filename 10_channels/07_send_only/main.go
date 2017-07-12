package main

import "fmt"

func ping(pings chan<- string, msg string) { // ping takes in send-only channel!
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // pong takes in a receive-only and a send-only channel!
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/*
The optinal chan<- marks the channel as send-only. Nothing can be received from this channel
*/

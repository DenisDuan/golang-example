package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message

	// By default sends and received block until both the sender and receiver are ready
	// This property allowed us to wait at the end of program for the "ping" message without
	// having to use any other synchronization.
	fmt.Println(msg)
}

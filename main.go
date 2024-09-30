package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	// Start a goroutine
	go func() {
		ch <- "Hello, World!"
	}()

	// Receive data from the channel
	msg := <-ch
	fmt.Println(msg) // Outputs: Hello, World!

}

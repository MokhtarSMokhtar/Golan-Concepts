package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

//Concurrency Patterns

//1- Generator Pattern
//A generator is a function that returns a channel.
//It encapsulates the creation of a goroutine and the communication channel.

func generatorWithBoring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}

func testGeneratorWithBoring() {
	c := generatorWithBoring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

//Fan-In Pattern
//Fan-In combines multiple channels into a single channel.
//Allows you to receive from multiple goroutines and handle their messages on a single channel.

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(generatorWithBoring("Joe"), generatorWithBoring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

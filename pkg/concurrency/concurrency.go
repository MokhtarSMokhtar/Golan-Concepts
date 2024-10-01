package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Concurrency in Go
// 1- Concurrency vs. Parallelism:
/*
Concurrency vs. Parallelism:
Concurrency is about structuring a program as the composition of independently executing processes.
Parallelism is about executing multiple processes simultaneously.
Go focuses on concurrency as a way to structure software that interacts with the real world, which is inherently concurrent.
*/
//Why Concurrency?
//To model real-world scenarios where multiple independent activities occur.
//Concurrency in Go simplifies programming by providing abstractions that are easy to reason about without dealing with low-level threading details.

// 2- Concurrency Concepts

// A- Goroutines
// To understand Goroutines let’s define a couple of terms.
// 1- Process: An instance of a program running on an operating system. It has its own memory space and resources.
//
// 2- Thread: A unit of execution within a process. Multiple threads can exist within a single process, sharing resources.
//
// 3- Goroutine: A lightweight thread managed by the Go runtime, allowing for concurrent execution of functions.
/*

	+-------------------------------+
	|           Operating System     |
	|                                |
	|  +-------------------------+   |
	|  |        Process A       |    |
	|  |                         |   |
	|  |  --------------------    |  |
	|  |  |   Thread 1       |    |  |
	|  |  |                  |    |  |
	|  |  +------------------+    |  |
	|  |  |   Thread 2       |    |  |
	|  |  +-----------------+    |   |
	|  |  |   Thread 3      |   |    |
	|  |  +-----------------+    |   |
	|  |                         |   |
	|  |  +-----------------+    |   |
	|  |  |   Goroutine 1   |   |    |
	|  |  |   (Lightweight) |   |    |
	|  |  +-----------------+    |   |
	|  |  |   Goroutine 2   |   |    |
	|  |  |   (Lightweight) |   |    |
	|  |  +-----------------+    |   |
	|  +-------------------------+   |
	|                                |
	|  +-------------------------+   |
	|  |        Process B       |    |
	|  |                         |   |
	|  |  +-----------------+    |   |
	|  |  |   Thread 1     |    |    |
	|  |  |                 |    |   |
	|  |  +-----------------+    |   |
	|  |                         |   |
	|  +-------------------------+   |
	+-------------------------------+

*/
//1.Definition:
//A goroutine is a lightweight thread managed by the Go runtime.
//It is an independently executing function launched using the go keyword.
//Goroutines have their own call stacks, which grow and shrink as needed.
//They are cheap to create, allowing for thousands or even millions to run concurrently.
//Initiation: When a Go program starts, the Go runtime creates a set of threads and launches a goroutine to execute the program.

//2. Goroutine Management
//Scheduler: The Go runtime has its own scheduler that automatically assigns goroutines to threads, similar to how an operating system schedules threads for CPU execution.

// 3.Syntax:
// go functionName()
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
func CallInMain() {
	go boring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

//In the example, boring runs as a goroutine, printing messages.
//The main function continues executing without waiting for boring to finish.
//When main returns, the program exits, even if goroutines are still running.
//--------------------------------------------------------------------------------------------------------------------------------
/// B- Channels
//In Go, a channel is a conduit through which goroutines communicate with each other.
//They allow you to send and receive values of a specified type between goroutines.
//Channels ensure synchronized communication.
//They are typed, meaning they can only transfer data of a specific type.
//Channels can be thought of as pipes connecting concurrent goroutines.

//Analogy: Imagine channels as conveyor belts in a factory where workers (goroutines) put and take items (data) in a synchronized manner.
/*

   +-------------------+          +-------------------+
   |    Worker A      |          |    Worker B      |
   | (Goroutine 1)    |          | (Goroutine 2)    |
   +-------------------+          +-------------------+
             |                               |
             |                               |
             V                               V
       +--------------------------------------------------+
       |                   Conveyor Belt                  |
       |            (Channel - Data Transfer)             |
       +--------------------------------------------------+
             ^                               ^
             |                               |
             |                               |
   +-------------------+          +-------------------+
   |    Worker C      |          |    Worker D      |
   | (Goroutine 3)    |          | (Goroutine 4)    |
   +-------------------+          +-------------------+

*/

// 1-  Creating Channels
// channelVariable := make(chan ValueType)
func createChannel() {
	ch := make(chan int)
	fmt.Printf("Channel: %v\n", ch)
}

// 2- channels are reference types. When you pass a channel to a function, you
//are really passing a pointer to the channel. Also like maps and slices, the zero value
//for a channel is nil.

// 3- Sending and Receiving Data
// Channels support two main operations:
// Sending data: Using the <- operator to the left of the channel.  ch <- value
// Receiving data: Using the <- operator to the right of the channel. value := <-ch
func sendAndReceivingData() {
	ch := make(chan string)
	// Start a goroutine
	go func() {
		ch <- "Hello, World!"
	}()

	// Receive data from the channel
	msg := <-ch
	fmt.Println(msg) // Outputs: Hello, World!
}

//Important Note:
//
//Unbuffered channels block the sending goroutine until another goroutine receives from the channel, and vice versa.

// 4-Buffered vs. Unbuffered Channels
// A-Unbuffered Channels
//No capacity to hold values; synchronization happens on send and receive.
//Both sender and receiver must be ready to proceed.
//ch := make(chan int) // Unbuffered channel

// B- Buffered Channels
//Have a capacity to hold a certain number of values.
//Send operations block only when the buffer is full.
//Receive operations block when the buffer is empty.
//ch := make(chan ValueType, capacity)

// 5- Closing Channels
//Use the close() function to close a channel.
//After closing, no more values can be sent.
//Receivers can still receive remaining buffered values.
//Attempting to send on a closed channel causes a panic.

func testCloseChannels() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	//The range loop continues to receive values from the channel until it's closed.
	//When the channel is closed and drained, the loop exits
	for val := range ch {
		fmt.Println(val)
	}
}

// 6- Select Statement
//The select statement allows a goroutine to wait on multiple communication operations.
/*
select {
case val := <-ch1:
    // Handle val from ch1
case ch2 <- val:
    // Send val to ch2
default:
    // Default case if no other case is ready
}
*/
func testSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()
	//The select statement waits until one of the cases is ready.
	//It proceeds with the case that is ready first.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

//WaitGroups
//A WaitGroup is a concurrency primitive that waits for a collection of goroutines to finish.
//It provides a way to block the execution of the main goroutine until other goroutines have completed their tasks.
//Package: sync
//Type: sync.WaitGroup

//When you use
//When you launch multiple goroutines and need to wait for all of them to complete before proceeding.
//Useful in scenarios where the main function should not exit before background tasks are done.

// 1- Using WaitGroups
//Add: Increment the WaitGroup counter for each goroutine you plan to wait for.
//Done: Each goroutine calls Done() when it finishes executing, which decrements the WaitGroup counter.
//Wait: The main goroutine calls Wait(), which blocks until the WaitGroup counter is zero.

func useWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1) // Increments the counter by 1
	wg.Done() // Decrements the counter by 1
	wg.Wait() // Blocks until the counter is 0

}

// Examples
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //defer get call whenever the func end
	fmt.Printf("Worker %d starting\n", id)
	// Simulate work
	// time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func testWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers completed.")
	//Worker 1 starting
	//Worker 2 starting
	//Worker 3 starting
	//Worker 4 starting
	//Worker 5 starting
	//Worker 2 done
	//Worker 1 done
	//Worker 3 done
	//Worker 4 done
	//Worker 5 done
	//All workers completed.

	//We create a WaitGroup called wg.
	//Before starting each goroutine, we call wg.Add(1) to increment the counter.
	//Each goroutine calls wg.Done() when it finishes.
	//The main goroutine calls wg.Wait() to block until the counter is zero.
}

// WaitGroup with HTTP Requests
func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	fmt.Printf("Fetched %s: %s\n", url, resp.Status)
	resp.Body.Close()
}
func callHttp() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("All URLs fetched.")
}

//Source :https://www.youtube.com/watch?v=f6kdp27TYZs&t=1739s

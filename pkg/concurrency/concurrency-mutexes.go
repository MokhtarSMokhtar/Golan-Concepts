// concurrency.go
//
// This package demonstrates the use of Mutexes in Go to manage
// concurrent access to shared resources, preventing race conditions.

package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ==============================
// Mutexes in Go
// ==============================
//
// Mutexes are synchronization primitives provided by Go's `sync` package.
// They are essential for controlling access to shared resources in concurrent programs,
// ensuring that only one goroutine can access a critical section of code at a time.
//
// By using mutexes, you can prevent race conditions and maintain data integrity
// when multiple goroutines interact with shared data.

// ==============================
// Types of Mutexes
// ==============================
//
// Go provides two primary types of mutexes:
//
// 1. sync.Mutex
//    A standard mutex that allows only one goroutine to access the critical section at a time.
//    It provides exclusive locking.
//
//    Declaration:
//        var mu sync.Mutex
//
// 2. sync.RWMutex
//    A Read-Write Mutex that allows multiple readers or one writer at a time.
//    It's useful when you have data that is read frequently but written infrequently.
//
//    Declaration:
//        var rwMu sync.RWMutex

// ==============================
// Using Mutexes
// ==============================
//
// To protect shared resources, follow these steps:
//
// 1. Lock the mutex before accessing the shared resource.
// 2. Perform the critical section operations.
// 3. Unlock the mutex after the operations are complete.
//
// This ensures that only one goroutine can execute the critical section at a time,
// preventing concurrent access issues.

// ==============================
// Examples
// ==============================

// Shared counter without mutex protection.
// Multiple goroutines increment this counter, leading to race conditions.
var counter int

// workerWithoutMutex increments the shared counter without any synchronization.
// This function demonstrates the potential for race conditions.
func workerWithoutMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

// Shared counter with mutex protection.
// This counter is safely incremented using a mutex to prevent race conditions.
var (
	counterWithMutex int
	mu               sync.Mutex
)

// workerWithMutex increments the shared counter with mutex protection.
// It locks the mutex before incrementing and unlocks it afterward.
func workerWithMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counterWithMutex++
		mu.Unlock()
	}
}

// SafeCounter is a thread-safe counter using mutex for synchronization.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely retrieves the current counter value.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// TestMutex demonstrates the difference between using mutexes and not using them.
// It runs multiple goroutines that increment shared counters and prints the final results.
func TestMutex() {
	var wg sync.WaitGroup

	fmt.Println("=== Mutex Demonstration ===")

	// Launch 5 goroutines without mutex protection
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workerWithoutMutex(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter (Without Mutex):", counter)
	// Explanation:
	// Expected counter is 5000 (5 workers * 1000 increments).
	// Due to race conditions, the final counter is likely less than 5000.

	// Reset counter for mutex-protected example
	counter = 0

	// Launch 5 goroutines with mutex protection
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workerWithMutex(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter (With Mutex):", counterWithMutex)
	// Explanation:
	// The mutex `mu` ensures that only one goroutine increments the counter at a time.
	// This prevents race conditions, resulting in the correct final counter value of 5000.

	// ==============================
	// SafeCounter Example
	// ==============================

	safeCounter := SafeCounter{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safeCounter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final SafeCounter:", safeCounter.Value())
	// Expected Output: Final SafeCounter: 5000
}

// ==============================
// Read-Write Mutex Example
// ==============================

// ReadHeavyStruct demonstrates the use of sync.RWMutex for read-heavy scenarios.
type ReadHeavyStruct struct {
	mu   sync.RWMutex
	data map[string]string
}

// Read retrieves data with a read lock.
func (r *ReadHeavyStruct) Read(key string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, exists := r.data[key]
	return val, exists
}

// Write modifies data with a write lock.
func (r *ReadHeavyStruct) Write(key string, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

// TestRWMutex demonstrates the usage of sync.RWMutex.
func TestRWMutex() {
	rhs := ReadHeavyStruct{
		data: make(map[string]string),
	}
	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rhs.Write("foo", fmt.Sprintf("value%d", i))
			fmt.Printf("Writer updated 'foo' to %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				val, ok := rhs.Read("foo")
				if ok {
					fmt.Printf("Reader %d read 'foo': %s\n", id, val)
				} else {
					fmt.Printf("Reader %d read 'foo': not found\n", id)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Map:", rhs.data)
}

// ==============================
// Atomic Operations Example
// ==============================

// atomicCounter demonstrates the use of atomic operations for a simple counter.
var atomicCounter int64

// atomicWorker safely increments the atomic counter.
func atomicWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&atomicCounter, 1)
	}
}

// TestAtomicCounter demonstrates using atomic operations instead of mutexes for simple synchronization.
func TestAtomicCounter() {
	var wg sync.WaitGroup

	// Launch 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go atomicWorker(&wg)
	}

	wg.Wait()
	fmt.Println("Final Atomic Counter:", atomicCounter)
	// Expected Output: Final Atomic Counter: 5000
}

// ==============================
// Deadlock Example
// ==============================

// deadlockExample demonstrates a scenario that leads to a deadlock.
func deadlockExample() {
	var mu sync.Mutex
	mu.Lock()
	// Forgot to unlock mu

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		// This goroutine will block forever waiting for mu
	}()
	wg.Wait()
	// The program will deadlock here
}

// ==============================
// Proper Locking Example
// ==============================

// properLockingExample demonstrates correct usage of mutexes to avoid deadlocks.
func properLockingExample() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		fmt.Println("Goroutine 1: Acquired lock")
		time.Sleep(1 * time.Second) // Simulate work
		mu.Unlock()
		fmt.Println("Goroutine 1: Released lock")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // Ensure Goroutine 1 locks first
		mu.Lock()
		fmt.Println("Goroutine 2: Acquired lock")
		mu.Unlock()
		fmt.Println("Goroutine 2: Released lock")
	}()

	wg.Wait()
}

// ==============================
// SafeQueue Example
// ==============================

// SafeQueue is a thread-safe queue implemented using a mutex.
type SafeQueue struct {
	mu    sync.Mutex
	queue []int
}

// Enqueue adds an item to the queue.
func (sq *SafeQueue) Enqueue(item int) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.queue = append(sq.queue, item)
}

// Dequeue removes and returns an item from the queue.
// Returns false if the queue is empty.
func (sq *SafeQueue) Dequeue() (int, bool) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	if len(sq.queue) == 0 {
		return 0, false
	}
	item := sq.queue[0]
	sq.queue = sq.queue[1:]
	return item, true
}

// producer adds items to the SafeQueue.
func producer(sq *SafeQueue, wg *sync.WaitGroup, items []int) {
	defer wg.Done()
	for _, item := range items {
		sq.Enqueue(item)
		fmt.Printf("Produced: %d\n", item)
	}
}

// consumer removes items from the SafeQueue.
func consumer(sq *SafeQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		item, ok := sq.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("Consumed: %d\n", item)
	}
}

// TestSafeQueue demonstrates using mutexes with channels and WaitGroups.
func TestSafeQueue() {
	sq := SafeQueue{}
	var wg sync.WaitGroup

	// Producer goroutine
	wg.Add(1)
	go producer(&sq, &wg, []int{1, 2, 3, 4, 5})

	// Consumer goroutine
	wg.Add(1)
	go consumer(&sq, &wg)

	wg.Wait()
	fmt.Println("Producer and Consumer completed.")
}

// ==============================
// Deadlock Prevention Example
// ==============================

// TestDeadlockPrevention demonstrates proper locking to avoid deadlocks.
func TestDeadlockPrevention() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		fmt.Println("Goroutine 1: Acquired lock")
		time.Sleep(1 * time.Second) // Simulate work
		mu.Unlock()
		fmt.Println("Goroutine 1: Released lock")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // Ensure Goroutine 1 locks first
		mu.Lock()
		fmt.Println("Goroutine 2: Acquired lock")
		mu.Unlock()
		fmt.Println("Goroutine 2: Released lock")
	}()

	wg.Wait()
}

// ==============================
// Main Function
// ==============================
//
// The main function serves as the entry point for the program.
// It calls various test functions to demonstrate mutex usage.

func main() {
	TestMutex()
	fmt.Println()

	TestRWMutex()
	fmt.Println()

	TestAtomicCounter()
	fmt.Println()

	// Uncomment the following line to see a deadlock (program will hang)
	// deadlockExample()

	properLockingExample()
	fmt.Println()

	TestSafeQueue()
	fmt.Println()

	// Uncomment the following line to see deadlock prevention
	// TestDeadlockPrevention()
}

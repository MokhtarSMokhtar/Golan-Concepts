// functions.go
//
// This package demonstrates the use of various types of functions in Go,
// including their syntax, return types, variadic functions, anonymous functions,
// closures, first-class functions, and the use of the `defer` keyword.
// It provides practical examples to illustrate each concept.

package functions

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// =============================
// Functions in Go
// =============================
//
// Functions are fundamental building blocks in Go, enabling you to encapsulate
// reusable code, organize your programs, and implement complex logic in a manageable way.
// Understanding how to define, use, and manipulate functions is essential for effective Go programming.
// This guide provides a comprehensive overview of functions in Go, covering their syntax,
// features, and practical examples.

// =============================
// 1. Syntax
// =============================
//
// The basic syntax for defining a function in Go is as follows:
//
//     func functionName(parameters) (returnTypes) {
//         // Function body
//     }
//
// Example:
// This function, `sub`, takes two integers and returns their difference.

func sub(x int, y int) int {
	return x - y
}

// =============================
// 2. Return Types
// =============================
//
// A function can return one or more values. The return type(s) are specified after the parameter list.
//
// Example:
// This function divides two floats and returns the result and an int.
// If division by zero is attempted, it returns predefined values.

func divide(a, b float64) (float64, int) {
	if b == 0 {
		return 1.3, 1 // Return float64 and int when division by zero
	}
	return a / b, 2
}

// =============================
// 3. Multiple Return Values
// =============================
//
// Go supports functions that return multiple values, which is particularly useful
// for error handling and returning additional information.
//
// Example:
// This function swaps two strings and returns them in reversed order.

func swap(a, b string) (string, string) {
	return b, a // Returns b, a
}

// Example Usage:
// Calling swap("hello", "world") returns "world", "hello".

// =============================
// 4. Named Return Values
// =============================
//
// You can name the return values in a function signature. Named return values act as
// variables defined at the top of the function and can be returned implicitly using
// the `return` statement without arguments.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Implicitly returns x and y
}

// =============================
// 5. Variadic Functions
// =============================
//
// Variadic functions are functions that accept a variable number of arguments.
// The type of the last parameter in a function definition is prefixed by ellipsis `...`,
// allowing the function to accept any number of arguments for that parameter.
//
// Example:
// This function calculates the sum of a variable number of integers.

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Example Usage:
// sum(1, 2, 3) // Returns 6
// sum(10, 20)  // Returns 30

// =============================
// 6. Anonymous Functions
// =============================
//
// Anonymous functions are functions without a name. They are often used for short-lived
// operations or as arguments to other functions.
//
// Example:
// This function demonstrates an anonymous function assigned to a variable.

func anonymousExample() {
	// Anonymous function assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("Multiply:", multiply(4, 5)) // Outputs: Multiply: 20
}

// =============================
// 7. Closures
// =============================
//
// A closure is an anonymous function that captures variables from its surrounding scope.
//
// Example:
// This function returns an adder function that adds a specific number to its argument.

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Example Usage:
// addFive := makeAdder(5)
// fmt.Println("Add Five:", addFive(3))  // Outputs: Add Five: 8
// fmt.Println("Add Five:", addFive(10)) // Outputs: Add Five: 15

// =============================
// 8. First-Class Functions
// =============================
//
// In Go, functions are first-class citizens, meaning they can be:
// 1. Assigned to variables.
// 2. Passed as arguments to other functions.
// 3. Returned from functions.
//
// Example:
// This function greets a person by name.

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// This function calls another function with a name.

func callFunction(f func(string), name string) {
	f(name)
}

// This function applies an operation on two integers.

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Example Operation Function: Adds two integers.

func add(a, b int) int {
	return a + b
}

// This function demonstrates first-class functions by passing functions as arguments.

func firstClassFunctionExample() {
	// Using callFunction to greet "Alice"
	callFunction(greet, "Alice") // Outputs: Hello, Alice!

	// Applying the add operation to 10 and 5
	resultOp := applyOperation(10, 5, add)
	fmt.Println("Apply Operation:", resultOp) // Outputs: Apply Operation: 15
}

// =============================
// 9. Ignoring Values
// =============================
//
// In Go, functions can return multiple values. However, there are situations where
// you might not need all the returned values. To handle such cases gracefully,
// Go provides a mechanism to ignore unwanted values using the blank identifier `_`.
//
// Example:
func getCoordinates() (int, int, int) {
	return 10, 20, 30
}

func ignoringValues() {
	// strconv.Atoi returns two values: an int and an error
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number:", num) // Outputs: Number: 123
	}
	// If you are certain that the string is a valid integer and want to ignore the error
	num, _ = strconv.Atoi("456")
	fmt.Println("Number without error check:", num) // Outputs: Number without error check: 456
	x, _, z := getCoordinates()
	fmt.Println("X:", x, "Z:", z) // Outputs: X: 10 Z: 30
}

// =============================
// 10. The `defer` Keyword
// =============================
//
// The `defer` keyword in Go schedules a function call to be executed after the
// surrounding function completes. This is particularly useful for resource cleanup tasks,
// such as closing files, unlocking mutexes, or releasing other resources,
// ensuring that these tasks are performed regardless of how the function exits
// (whether normally or due to an error).

func readFile(dstName, srcName string) {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Perform file operations here
	fmt.Printf("Reading from %s to %s\n", srcName, dstName)
	// Additional file handling logic
}

// =============================
// Important Topics
// =============================
//
// This section highlights crucial aspects of functions in Go that are essential
// for writing clean, efficient, and bug-free code.

// =============================
// 11. Embedding Mutexes in Structs
// =============================
//
// Embedding mutexes within structs is a common practice to protect the struct's fields
// from concurrent access. This approach encapsulates the synchronization mechanism,
// promoting better code organization and safety.

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment increases the counter by one.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// TestSafeCounter demonstrates the usage of SafeCounter with mutex protection.

func TestSafeCounter() {
	var wg sync.WaitGroup
	counter := SafeCounter{}

	// Launch 5 goroutines to increment the counter.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter (SafeCounter):", counter.Value())
	// Expected Output: Final Counter (SafeCounter): 5000
}

// =============================
// 12. Performance Considerations: `sync.Mutex` vs. `sync.RWMutex`
// =============================
//
// Choosing between `sync.Mutex` and `sync.RWMutex` depends on the nature of your application's
// read and write operations. Understanding their performance implications can help optimize
// your concurrent programs.

//
// | Feature                | `sync.Mutex`                                | `sync.RWMutex`                                   |
// |------------------------|---------------------------------------------|--------------------------------------------------|
// | **Lock Type**          | Exclusive (Write)                           | Read and Write                                   |
// | **Read Concurrency**  | No, exclusive access                       | Yes, multiple readers can hold the lock simultaneously |
// | **Write Concurrency** | No, exclusive access                       | No, only one writer can hold the lock               |
// | **Use Case**           | When both reads and writes are frequent     | When reads are much more frequent than writes      |
// | **Performance**        | Generally faster for write-heavy workloads  | Better performance for read-heavy workloads       |
//

// =============================
// 13. Comprehensive Example
// =============================
//
// This section combines multiple concepts to demonstrate a comprehensive use case
// involving functions, goroutines, mutexes, and defer statements.

func comprehensiveExample() {
	var wg sync.WaitGroup

	// Initialize SafeCounter
	counter := SafeCounter{}

	// Launch 5 goroutines to increment the counter concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d finished incrementing.\n", id)
		}(i + 1)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter (Comprehensive Example):", counter.Value())
}

// =============================
// Main Function
// =============================
//
// The main function serves as the entry point for the program.
// It demonstrates the usage of various functions defined above.

func main() {
	fmt.Println("=== Functions in Go ===")

	// 1. Basic Function Usage
	fmt.Println("Subtraction:", sub(10, 5)) // Outputs: Subtraction: 5

	// 2. Function with Multiple Return Values
	result, code := divide(10, 2)
	fmt.Println("Division:", result, "Code:", code) // Outputs: Division: 5 Code: 2

	// Attempting to divide by zero
	result, code = divide(10, 0)
	fmt.Println("Division:", result, "Code:", code) // Outputs: Division: 1.3 Code: 1

	// 3. Swapping Values
	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b) // Outputs: Swap: world hello

	// 4. Named Return Values
	x, y := split(17)
	fmt.Println("Split:", x, y) // Outputs: Split: 8 9

	// 5. Variadic Function Usage
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5)) // Outputs: Sum: 15
	fmt.Println("Sum:", sum(10, 20))        // Outputs: Sum: 30

	// 6. Anonymous Function Example
	anonymousExample() // Outputs: Multiply: 20

	// 8. First-Class Functions Example
	firstClassFunctionExample()

	// 9. Ignoring Values Example
	ignoringValues()

	// 10. Defer Keyword Example
	readFile("destination.txt", "source.txt")

	// 11. SafeCounter with Mutex Example
	TestSafeCounter()

	// 12. Comprehensive Example Combining Multiple Concepts
	comprehensiveExample()

	fmt.Println("=== End of Functions Demonstration ===")
}

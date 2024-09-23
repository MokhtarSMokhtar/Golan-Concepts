package main

import (
	"fmt"
)

// Golan Functions
/*
Functions enable you to encapsulate reusable code,
organize your programs, and implement complex logic in a manageable way.
Understanding how to define, use, and manipulate functions is essential for effective Go programming.
This guide provides a comprehensive overview of functions in Go,
covering their syntax, features, and practical examples.
Functions in Go can take zero or more arguments.
To make code easier to read, the variable type comes after the variable name.
*/

// 1- Syntax

/*
// Basic Function Syntax
func functionName(parameters) (returnTypes) {
    // Function body
}
*/

// This function, sub, takes two integers and returns their difference.
func sub(x int, y int) int {
	return x - y
}

// 2- Return Types

/*
A function can return one or more values.
The return type(s) are specified after the parameter list.
*/

// This function divides two floats and returns the result and an int.
// If division by zero is attempted, it returns a predefined float64 and int.
func divide(a, b float64) (float64, int) {
	if b == 0 {
		return 1.3, 1 // Return float64, int when division by zero
	}
	return a / b, 2
}

// 3- Multiple Return Values

/*
Go supports functions that return multiple values,
which is particularly useful for error handling and returning additional information.
*/

// This function swaps two strings and returns them in reversed order.
func swap(a, b string) (string, string) {
	return b, a // Return b, a
}

// Example:
// Calling swap("hello", "world") returns "world", "hello".

// 4- Named Return Values

/*
You can name the return values in a function signature.
Named return values act as variables defined at the top of the function
and can be returned implicitly using the return statement without arguments.
*/

// This function splits a sum into two parts and returns them.
// Here, x and y are named return values.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Implicitly returns x and y
}

////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////Important Topics//////////////////////////////////////////
//

// 1- Variadic Functions

/*
Usually, functions in Go accept only a fixed number of arguments.
However, it is also possible to write variadic functions in Go.
A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis ...,
then the function can accept any number of arguments for that parameter.

Example:
func find(a int, b ...int) {
    // ...
}
*/

// This function calculates the sum of a variable number of integers.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// You can call sum with any number of integer arguments:
// sum(1, 2, 3) // Returns 6
// sum(10, 20)  // Returns 30

// 2- Anonymous Functions

/*
Anonymous functions are functions without a name.
They are often used for short-lived operations or as arguments to other functions.

Example:
func() {
    fmt.Println("Hello from anonymous function")
}()
*/

// This function demonstrates an anonymous function assigned to a variable.
func anonymousExample() {
	// Anonymous function assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("Multiply:", multiply(4, 5)) // Outputs: Multiply: 20
}

// 3- Closures

/*
A closure is an anonymous function that captures variables from its surrounding scope.

Example:
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}
*/

// This function returns an adder function that adds a specific number to its argument.
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// This function demonstrates the use of closures.
func exampleInMain() {
	addFive := makeAdder(5)
	fmt.Println("Add Five:", addFive(3))  // Outputs: Add Five: 8
	fmt.Println("Add Five:", addFive(10)) // Outputs: Add Five: 15
}

// 4- First-Class Functions

/*
In Go, functions are first-class citizens, meaning they can be:
1- Assigned to variables.
2- Passed as arguments to other functions.
3- Returned from functions.
*/

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

// Example operation function: adds two integers.
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

////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////Main Function/////////////////////////////////////////
//

// The main function demonstrates the usage of all the functions defined above.
func main() {
	// 1. Calling sub function
	fmt.Println("Subtraction:", sub(10, 5)) // Outputs: Subtraction: 5

	// 2. Calling divide function
	result, code := divide(10, 2)
	fmt.Println("Division:", result, "Code:", code) // Outputs: Division: 5 Code: 2

	// Attempting to divide by zero
	result, code = divide(10, 0)
	fmt.Println("Division:", result, "Code:", code) // Outputs: Division: 1.3 Code: 1

	// 3. Calling swap function
	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b) // Outputs: Swap: world hello

	// 4. Calling split function
	x, y := split(17)
	fmt.Println("Split:", x, y) // Outputs: Split: 8 9

	// 5. Calling sum function with multiple arguments
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5)) // Outputs: Sum: 15
	fmt.Println("Sum:", sum(10, 20))        // Outputs: Sum: 30

	// 6. Calling anonymousExample function
	anonymousExample() // Outputs: Multiply: 20

	// 7. Using closures
	exampleInMain()

	// 8. Using first-class functions
	firstClassFunctionExample()
}

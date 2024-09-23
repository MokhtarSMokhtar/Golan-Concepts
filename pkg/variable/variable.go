// package variable demonstrates basic variables in Go.
package variable

import "fmt" // Import fmt package for printing

/*
1. Variable Declaration
Go provides two primary ways to declare variables:
   a. Using the "var" Keyword
   b. Using Short Variable Declaration ":="
*/

// 1.a Using var
/*
   var variableName type
   var variableName type = value
   var variableName = value // Type inferred
*/

// Examples of variable declarations using var
var age int
var name string = "Your Name"
var isStudent = true

// 1.b Using Short Variable Declaration
/*
   variableName := value
   - Only usable inside functions.
   - Automatically infers the type based on the assigned value.
*/

// Example of short variable declaration (commented out to avoid errors outside functions)
// score := 95.5 // Uncommenting this line will cause an error if not inside a function

//////////////////////////////////////////////////////////////////////////////////////////////

/*
2. Type Inference
Go can automatically determine the type of a variable based on the value assigned to it.
*/

// Function demonstrating type inference
func typeInference() {
	var count = 10      // Type inferred as int
	flag := true        // Type inferred as bool
	temperature := 36.6 // Type inferred as float64
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
	fmt.Println(count, flag, temperature)
}

//////////////////////////////////////////////////////////////////////////////////////////////

/*
3. Zero Values
If a variable is declared without an explicit initial value, Go assigns it a zero value based on its type.
Common Zero Values:
   - int, float64: 0, 0.0
   - string: "" (empty string)
   - bool: false
   - Pointers, Slices, Maps, Channels, Interfaces: nil
*/

// Example of variables with zero values
var number int     // 0
var price float64  // 0.0
var message string // ""
var isActive bool  // false
var ptr *int       // nil

//////////////////////////////////////////////////////////////////////////////////////////////

/*
4. Constants
Constants are immutable values that are known at compile time
and do not change during the execution of the program.
   const constantName type = value
   const constantName = value // Type inferred
*/

// Examples of constants
const Greeting = "Welcome to Go!"
const MaxUsers = 100

//////////////////////////////////////////////////////////////////////////////////////////////

/*
5. Naming Conventions
   A. Exported Variables: Start with an uppercase letter (accessible from other packages).
      var ExportedVar int
   B. Unexported Variables: Start with a lowercase letter (accessible only within the package).
      var unexportedVar string
*/

//////////////////////////////////////////////////////////////////////////////////////////////

/*
6. Variable Scope
   - Global Variables: Declared outside of functions, accessible throughout the package.
*/

// Example of a global variable
var globalVar int = 100

func a() {
	// Can access globalVar here
	var localVar string = "Hello"
	fmt.Println(localVar) // Accessible here
}

func anotherFunction() {
	// fmt.Println(localVar) // Error: undefined localVar
}

/*
Closures
A closure is a function that references variables from outside its own function body.
The function may access and assign to the referenced variables.
(Ignore the function definition)
FOR MORE ABOUT lexical scope: https://stackoverflow.com/a/1047491
*/

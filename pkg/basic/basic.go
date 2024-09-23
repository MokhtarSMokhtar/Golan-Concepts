package basic

// What's Go?
/*
Golang is an open-source programming language designed by Google.
It was created in 2007 by Robert Griesemer, Rob Pike,
and Ken Thompson to address the limitations of existing languages,
particularly in terms of simplicity, speed, and scalability.
*/

//**************************
// Concepts of Golang
/*
1. Expression-oriented:
   Golang is an expression-oriented language,
   which means that expressions can be used in a variety of contexts,
   including as statements, return values, and arguments to functions.

   Example:
   func add(a, b int) int {
       return a + b // The expression a + b is evaluated and returned.
   }

2. Static typing:
   Golang is a statically typed language,
   which means that the type of each variable must be declared before it is used.
   This can help to prevent errors at runtime.

   Example:
   var x int = 10 // x is explicitly declared as an int.

3. Concurrency:
   Go provides built-in support for concurrent programming through goroutines and channels.
   This allows developers to write programs that can handle multiple tasks simultaneously.

   Example:
   go func() {
       // Code to run concurrently
   }()

4. Garbage Collection:
   Go has an automatic garbage collector, which helps manage memory automatically.
   This reduces memory leaks and allows developers to focus on writing code without worrying too much about memory management.

5. Simple Syntax:
   Go's syntax is designed to be clean and easy to understand,
   making it accessible to beginners while still being powerful for experienced developers.

6. Package Management:
   Go has a robust package management system,
   making it easy to manage dependencies and share code through modules.

7. Strong Standard Library:
   Go includes a rich standard library that provides a wide range of functionality,
   including support for networking, cryptography, and web development.

8. Package-based:
   Golang is a package-based language, which means that code is organized into packages.
   This can help to improve code modularity and reusability.

9. Pointers and Interfaces:
   Golang supports pointers and interfaces,
   which are powerful tools for memory management and object-oriented programming.

10. First-Class Functions:
    Functions in Go are first-class citizens, meaning they can be assigned to variables,
    passed as arguments, and returned from other functions.


11. Error Handling:
    Go uses a simple error handling approach based on multiple return values.
    Functions often return an error value alongside the main return value.


12. Structs and Composition:
    Go uses structs to define complex data types and supports composition,
    allowing developers to build complex types using simpler types.

12. Structs and Composition:
    Go uses structs to define complex data types and supports composition,
    allowing developers to build complex types using simpler types.

13. Fast and Compiled::
Go is a compiled language, meaning it generates machine code,
making its programs faster than interpreted ones like JavaScript,
Python, or Ruby. While it may not match the speed of compiled languages like Rust, Zig, or C,
Go's strength lies in its much faster compilation times,
leading to a more productive developer experienc

*/

// Example Function

package arrays

import "fmt"

//Array in Golan
/*
Arrays in Go are fixed-size collections of elements of the same type,
with their size determined at compile time and considered part of their type. They are value types,
so assigning one array to another or passing them to functions copies all their elements. While arrays are less commonly used due to their rigidity,
they are fundamental to understanding slices, which offer dynamic sizing and are more prevalent in Go programming.
*/

// 1- Declaration
/*
var arrayName [size]elementType
*/
var arr [5]int

// 2- Initializing Arrays
var numbers = [5]int{1, 2, 3, 4, 5}

// Letting the compiler determine the array size:
var numbersDy = [...]int{1, 2, 3, 4, 5}

// The compiler infers the size as 5

// /Initializing specific elements:
func init() {
	// Only initialize certain indices
	numbers := [5]int{2: 100, 4: 200}
	fmt.Println(numbers) // Outputs: [0 0 100 0 200]

}

//Arrays as Value Types
/*
Arrays are value types in Go. When you assign an array to another variable or pass it to a function, a copy of the array is made
so if you will pass the array to function you should pass it by reference
*/

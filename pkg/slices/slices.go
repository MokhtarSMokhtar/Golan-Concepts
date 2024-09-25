package slices

import "fmt"

// Slice in Go
/*
A slice is a dynamically-sized, flexible view into the elements of an array. Unlike arrays,
slices are reference types,
and they provide a powerful way to manage collections of data without the rigidity of fixed-size arrays.
Key Characteristics of Slices:
Dynamic size: Can grow and shrink as needed.
Reference type: Points to an underlying array.
Efficient: Only the slice descriptor is copied when passing slices to functions.
*/

//Difference Between Arrays and Slices
/*
Arrays:

Fixed size: The size of an array is part of its type.
Value type: Assigning one array to another copies all elements.
Less commonly used directly due to rigidity.
Slices:

Dynamic size: Can change length during execution.
Reference type: Points to an underlying array.
Preferred for working with sequences of data.
*/

//Example
//var arr [5]int     // Array of fixed size 5
//var slc []int      // Slice of ints (no fixed size)

// 1- Creating Slices

// // A-From an Array
func createSliceFromArray() {
	//syntax slice := array[startIndex:endIndex]
	arr := [5]int{1, 2, 3, 4, 5}
	slc := arr[1:4]  // Creates a slice from arr[1] to arr[3]
	fmt.Println(slc) // Outputs: [2 3 4]
}

// / B- using make function
// The make function is used to create slices with a specified length and capacity.
func createSliceWithMake() {
	//syntax slice := make([]int, length, capacity)
	slc := make([]int, 5, 10)
	fmt.Println("Length:", len(slc))   // Outputs: Length: 5
	fmt.Println("Capacity:", cap(slc)) // Outputs: Capacity: 10
}

// / C- Slice Literals
var slice = []int{1, 2, 3, 4, 5}

// 2- Nil and Empty Slices
// Nil Slice: A slice with no underlying array.
func nilSlice() {

	var slc []int
	fmt.Println(slc == nil) // Outputs: true
	//empty slice
	emptySlc := []int{}
	fmt.Println(len(emptySlc))   // Outputs: 0
	fmt.Println(emptySlc == nil) // Outputs: false

}

// 3- Manipulating Slices

// / A-Accessing Elements
func useSliceIndex() {
	//
	slice := []int{10, 20, 30}
	fmt.Println(slice[0])
}

// / B-Appending Elements
func appendToSlice() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice) // Outputs: [1 2 3 4 5]
}

// 4- Capacity and Length
//Understanding len and cap is essential for efficient slice management.
/// A- len
/*he number of elements in the slice.
Accessible via the built-in len function.
slice := []int{1, 2, 3}
fmt.Println(len(slice)) // Outputs: 3
*/

/// B- Capacity (cap)
/*
The maximum length the slice can reach "without reallocation."
Accessible via the built-in cap function
slice := make([]int, 3, 5)
fmt.Println(cap(slice)) // Ou
*/

// 5- Iterating Over Slices
//Use the for loop or for range loop to iterate over slices.

// /A- for
func useForLoop() {
	slice := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}

// /B- for range
func useForRange() {
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	/*
		for _, value := range slice {
		    fmt.Println(value)
		}
	*/
}

// 6- Removing Elements
// Removing-elements from a slice can be tricky due to the underlying array.
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

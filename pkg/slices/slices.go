// slices.go
//
// This package demonstrates the use of Slices in Go, including their creation,
// manipulation, and key characteristics. It provides practical examples to
// illustrate various slice operations and best practices for effective slice usage.

package slices

import (
	"fmt"
)

// =============================
// Slices in Go
// =============================
//
// A slice is a dynamically-sized, flexible view into the elements of an array.
// Unlike arrays, slices are reference types and provide a powerful way to manage
// collections of data without the rigidity of fixed-size arrays.
//
// **Key Characteristics of Slices:**
// - **Dynamic Size:** Can grow and shrink as needed.
// - **Reference Type:** Points to an underlying array.
// - **Efficient:** Only the slice descriptor is copied when passing slices to functions.

// =============================
// 1. Difference Between Arrays and Slices
// =============================
//
// **Arrays:**
// - **Fixed Size:** The size of an array is part of its type.
// - **Value Type:** Assigning one array to another copies all elements.
// - **Less Flexible:** Due to fixed size, arrays are less commonly used directly.
//
// **Slices:**
// - **Dynamic Size:** Can change length during execution.
// - **Reference Type:** Points to an underlying array.
// - **Flexible and Efficient:** Preferred for working with sequences of data.

// Example Declarations:
// var arr [5]int     // Array of fixed size 5
// var slc []int      // Slice of ints (no fixed size)

// =============================
// 2. Creating Slices
// =============================
//
// Slices can be created in several ways:
// 1. From an Array
// 2. Using the `make` Function
// 3. Slice Literals

// 2A. Creating a Slice from an Array
func createSliceFromArray() {
	// Syntax: slice := array[startIndex:endIndex]
	arr := [5]int{1, 2, 3, 4, 5}
	slc := arr[1:4]                       // Creates a slice from arr[1] to arr[3]
	fmt.Println("Slice from array:", slc) // Outputs: [2 3 4]
}

// 2B. Creating a Slice using the `make` Function
func createSliceWithMake() {
	// Syntax: slice := make([]int, length, capacity)
	slc := make([]int, 5, 10)
	fmt.Println("Slice created with make:")
	fmt.Println("Slice:", slc)         // Outputs: [0 0 0 0 0]
	fmt.Println("Length:", len(slc))   // Outputs: 5
	fmt.Println("Capacity:", cap(slc)) // Outputs: 10
}

// 2C. Slice Literals
var sliceLiteral = []int{1, 2, 3, 4, 5} // Slice literal

// =============================
// 3. Nil and Empty Slices
// =============================
//
// - **Nil Slice:** A slice with no underlying array.
// - **Empty Slice:** A slice with zero length but not nil.

func nilAndEmptySlices() {
	var nilSlc []int
	fmt.Println("nilSlc is nil:", nilSlc == nil) // Outputs: true

	emptySlc := []int{}
	fmt.Println("emptySlc length:", len(emptySlc))   // Outputs: 0
	fmt.Println("emptySlc is nil:", emptySlc == nil) // Outputs: false
}

// =============================
// 4. Manipulating Slices
// =============================
//
// Slices can be manipulated in various ways, including accessing elements,
// appending elements, and more.

// 4A. Accessing Elements
func accessSliceElements() {
	slc := []int{10, 20, 30}
	fmt.Println("First element:", slc[0])  // Outputs: 10
	fmt.Println("Second element:", slc[1]) // Outputs: 20
	fmt.Println("Third element:", slc[2])  // Outputs: 30
}

// 4B. Appending Elements
func appendToSlice() {
	slc := []int{1, 2, 3}
	fmt.Println("Original Slice:", slc) // Outputs: [1 2 3]

	slc = append(slc, 4, 5)
	fmt.Println("Appended Slice:", slc) // Outputs: [1 2 3 4 5]
}

// =============================
// 5. Capacity and Length
// =============================
//
// Understanding `len` and `cap` is essential for efficient slice management.
//
// - **len(slice):** The number of elements in the slice.
// - **cap(slice):** The maximum number of elements the slice can hold without reallocating.

// Example usage in `createSliceWithMake` function above.

// =============================
// 6. Iterating Over Slices
// =============================
//
// You can iterate over slices using traditional `for` loops or the `for range` construct.

// 6A. Using a Traditional `for` Loop
func iterateWithFor() {
	slc := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(slc); i++ {
		fmt.Println("Index", i, ":", slc[i])
	}
	// Outputs:
	// Index 0 : apple
	// Index 1 : banana
	// Index 2 : cherry
}

// 6B. Using a `for range` Loop
func iterateWithForRange() {
	slc := []string{"apple", "banana", "cherry"}
	for index, value := range slc {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	// Outputs:
	// Index: 0, Value: apple
	// Index: 1, Value: banana
	// Index: 2, Value: cherry

	// If the index is not needed, use the blank identifier `_`
	for _, value := range slc {
		fmt.Println("Value:", value)
	}
	// Outputs:
	// Value: apple
	// Value: banana
	// Value: cherry
}

// =============================
// 7. Removing Elements
// =============================
//
// Removing elements from a slice can be done using the `append` function to
// exclude the element at a specific index. This operation modifies the slice
// by creating a new underlying array if necessary.

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

// =============================
// 8. Comprehensive Example
// =============================
//
// This section combines multiple slice operations to demonstrate a robust example.

func comprehensiveSliceExample() {
	fmt.Println("=== Comprehensive Slice Example ===")

	// 1. Creating slices
	createSliceFromArray()
	createSliceWithMake()
	fmt.Println("Slice Literal:", sliceLiteral) // Outputs: [1 2 3 4 5]

	// 2. Nil and Empty slices
	nilAndEmptySlices()

	// 3. Accessing and Appending
	accessSliceElements()
	appendToSlice()

	// 4. Iterating
	iterateWithFor()
	iterateWithForRange()

	// 5. Removing Elements
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slc) // Outputs: [1 2 3 4 5]
	slc = removeElement(slc, 2)
	fmt.Println("After Removing Index 2:", slc) // Outputs: [1 2 4 5]

	fmt.Println("=== End of Comprehensive Slice Example ===")
}

// =============================
// Main Function
// =============================
//
// The main function serves as the entry point for the program.
// It demonstrates the usage of various slice operations defined above.

func main() {
	fmt.Println("=== Slices in Go ===")

	// 1. Creating Slices
	createSliceFromArray()
	createSliceWithMake()
	fmt.Println("Slice Literal:", sliceLiteral) // Outputs: [1 2 3 4 5]

	// 2. Nil and Empty Slices
	nilAndEmptySlices()

	// 3. Manipulating Slices
	accessSliceElements()
	appendToSlice()

	// 4. Iterating Over Slices
	iterateWithFor()
	iterateWithForRange()

	// 5. Removing Elements
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slc) // Outputs: [1 2 3 4 5]
	slc = removeElement(slc, 2)
	fmt.Println("After Removing Index 2:", slc) // Outputs: [1 2 4 5]

	// 6. Comprehensive Example
	comprehensiveSliceExample()

	fmt.Println("=== End of Slices Demonstration ===")
}

package generics

import (
	"fmt"
	"sort"
)

// Package generics provides examples and explanations of generics in Go.
// It demonstrates how to use generics to create flexible and reusable
// functions, types, and data structures without sacrificing type safety.
//
// Generics were introduced in Go 1.18, enabling developers to write
// type-agnostic code that operates on different data types with
// compile-time type checking.

// ============================================================================
// Understanding Generics in Go
// ============================================================================

// What Are Generics?
//
// Generics allow you to write flexible and reusable functions, types, and
// data structures that can operate with different data types without
// sacrificing type safety. They enable you to abstract over types, making
// your code more generic and reducing duplication.
//
// Generics in Go: An Overview
//
// Go introduced generics in Go 1.18. With generics, you can define functions
// and data structures that work with any data type, specified when you use
// them. This is particularly useful for creating reusable components like
// data structures (e.g., linked lists, stacks) and utility functions (e.g.,
// sorting, searching).

// Key Concepts:
//
// 1. Type Parameters:
//    - Represented within square brackets `[]` in function and type definitions.
//    - Serve as placeholders for actual types specified during usage.
//
// 2. Type Constraints:
//    - Define the set of types that can be used as type arguments.
//    - Ensure that type parameters meet certain requirements (e.g.,
//      implement specific interfaces).

// ============================================================================
// Syntax and Examples
// ============================================================================

// 1. Generic Function
//
// Example: A generic function to find the maximum element in a slice.

// Comparable is a type constraint that allows any type that supports the > operator.
// The tilde `~` allows for underlying type definitions (e.g., custom types based on int).
type Comparable interface {
	~int | ~float64 | ~string
}

// Max returns the maximum element in a slice.
// It works with any type that satisfies the Comparable constraint.
func Max[T Comparable](slice []T) T {
	if len(slice) == 0 {
		panic("slice is empty")
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func ExampleMax() {
	ints := []int{1, 3, 2, 5, 4}
	floats := []float64{1.1, 3.3, 2.2, 5.5, 4.4}
	stringsSlice := []string{"apple", "orange", "banana", "pear"}

	fmt.Println("Max int:", Max(ints))            // Output: Max int: 5
	fmt.Println("Max float:", Max(floats))        // Output: Max float: 5.5
	fmt.Println("Max string:", Max(stringsSlice)) // Output: Max string: pear
}

// Explanation:
//
// Type Parameter `[T Comparable]`:
// - `T` is a type parameter constrained by the `Comparable` interface.
// - `Comparable` allows types that can be compared using the `>` operator (`int`, `float64`, `string`).
//
// Function `Max[T Comparable]`:
// - Accepts a slice of any type `T` that satisfies `Comparable`.
// - Returns the maximum element of type `T`.

// 2. Generic Data Structure
//
// Example: A generic Pair struct that holds two values of any types.

// Pair holds two values of any types.
type Pair[T, U any] struct {
	First  T
	Second U
}

func ExamplePair() {
	intStringPair := Pair[int, string]{First: 1, Second: "one"}
	stringBoolPair := Pair[string, bool]{First: "isActive", Second: true}

	fmt.Println(intStringPair)  // Output: {1 one}
	fmt.Println(stringBoolPair) // Output: {isActive true}
}

// Explanation:
//
// Type Parameters `[T, U any]`:
// - `T` and `U` are type parameters that can be any type (`any` is a predefined constraint that allows any type).
//
// Struct `Pair[T, U]`:
// - Holds two fields, `First` of type `T` and `Second` of type `U`.
//
// Usage:
// - Creates instances of `Pair` with different type arguments (`int, string` and `string, bool`).

// 3. Generic Interface and Constraints
//
// Example: A generic SortSlice function that sorts a slice of any ordered type.

// Ordered is a type constraint that allows any ordered type.
// It includes types that support the < operator, such as int, float64, and string.
type Ordered interface {
	~int | ~float64 | ~string
}

// SortSlice sorts a slice of any ordered type.
func SortSlice[T Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func ExampleSortSlice() {
	ints := []int{5, 3, 4, 1, 2}
	floats := []float64{5.5, 3.3, 4.4, 1.1, 2.2}
	stringsSlice := []string{"banana", "apple", "pear", "orange"}

	SortSlice(ints)
	SortSlice(floats)
	SortSlice(stringsSlice)

	fmt.Println("Sorted ints:", ints)            // Output: Sorted ints: [1 2 3 4 5]
	fmt.Println("Sorted floats:", floats)        // Output: Sorted floats: [1.1 2.2 3.3 4.4 5.5]
	fmt.Println("Sorted strings:", stringsSlice) // Output: Sorted strings: [apple banana orange pear]
}

// Explanation:
//
// Type Constraint `Ordered`:
// - Allows types that can be ordered (`int`, `float64`, `string`).
//
// Function `SortSlice[T Ordered]`:
// - Sorts a slice of any type `T` that satisfies the `Ordered` constraint using Go's `sort` package.
//
// Usage:
// - Sorts slices of `int`, `float64`, and `string`.

// ============================================================================
// Benefits of Using Generics
// ============================================================================

// Reusability:
// - Write functions and data structures once and use them with different types.
//
// Type Safety:
// - Retain compile-time type checking without resorting to interfaces and type assertions.
//
// Performance:
// - Avoids the overhead of reflection or boxing/unboxing since types are determined at compile time.
//
// Cleaner Code:
// - Reduces code duplication, making your codebase easier to maintain and read.

// ============================================================================
// When to Use Generics
// ============================================================================

// Reusable Data Structures:
// - Implementing generic collections like lists, stacks, queues, maps, etc.
//
// Utility Functions:
// - Functions that perform operations like sorting, filtering, mapping over data.
//
// Type-Agnostic Algorithms:
// - Algorithms that operate on data regardless of the underlying type, provided they meet certain constraints.

// ============================================================================
// Limitations
// ============================================================================

// Complexity:
// - Introduces additional complexity in understanding and maintaining type constraints.
//
// Compilation Time:
// - Can increase compilation time, especially with extensive use of generics.
//
// Overuse:
// - Using generics where simple type assertions or interfaces suffice can lead to unnecessarily complicated code.

package pointers

import "fmt"

//What a pointers ?
//A pointer is a variable that holds the memory address of another variable.
//Instead of holding a value directly, a pointer points to the location in memory where the value is stored.

//Why Use Pointers?
/// 1- Efficiency: Passing pointers to large data structures avoids copying the entire structure, saving memory and processing time.
/*
 When we have large amounts of data, making copies to pass between functions is very inefficient.
 By passing the memory location of where the data is stored instead, we can dramatically reduce the resource-footprint of our programs.
 By passing pointers between functions, we can access and modify the single copy of the data directly,
 meaning that any changes made by one function are immediately visible to other parts of the program when the function ends.
*/
/// 2- Data Structures: Pointers are essential for creating complex data structures like linked lists, trees, and graphs.
/// 3- Shared Access: Multiple parts of a program can access and modify the same data.

// /Pointer Basics
func explainBasicOfPointer() {
	///Declaration
	/////To declare a pointer variable, use the * operator before the type.
	//var prt* Type
	var num int = 42
	var ptr *int = &num
	//ptr is a pointer to an int.
	//&num gets the memory address of num.
	fmt.Println(ptr) // Outputs something like 0xc0000140b0

	///The * (Dereference) Operator
	//The * operator is used to:
	//
	//Declare a pointer type.
	//Dereference a pointer to access or modify the value it points to.
}

// Pointers and Functions
// /Passing by Value vs. Passing by Reference
// In Go, arguments are passed by value. When you pass a variable to a function, it creates a copy.
func increment(x int) {
	x = x + 1
}
func incrementWitPoi(x *int) {
	*x = *x + 1
}

func call() {
	num := 42
	increment(num)
	fmt.Println(num) // Outputs: 42 (num is unchanged)
	//To modify the original variable, pass a pointer.
	incrementWitPoi(&num)
	fmt.Println(num) // Outputs: 43 (num is modified)
}

///Pointers to Structs
///When working with structs, pointers allow you to modify the original struct or efficiently pass large structs without copying.
///Example

type Person struct {
	Name string
	Age  int
}

func updateName(p *Person, newName string) {
	p.Name = newName
}
func mainCall() {
	person := Person{Name: "Alice", Age: 30}
	updateName(&person, "Bob")
	fmt.Println(person.Name) // Outputs: Bob

	///Note: Go allows you to access struct fields through a pointer without explicit dereferencing.
	//p := &Person{Name: "Alice", Age: 30}
	//p.Age = 31 // Equivalent to (*p).Age = 31
}

// The new Function
// The new function allocates memory for a type and returns a pointer to it.
func callNewFunc() {
	ptr := new(int) // ptr is of type *int
	*ptr = 42
	fmt.Println(*ptr) // Outputs: 42
	//new(Type) returns a pointer to a zero-initialized value of the specified type.
	//It's equivalent to:
	//var ptr *int = &int{}
}

//Pointer Types and Zero Values
//The zero value of a pointer type is nil.

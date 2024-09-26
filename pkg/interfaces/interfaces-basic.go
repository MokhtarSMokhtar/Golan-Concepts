package interfaces

import (
	"fmt"
	"io"
)

//Interfaces in Golang
/*
Interfaces are a fundamental and powerful feature in Go that enable polymorphic behavior and decoupling between different parts of your code.
They allow you to define a set of method signatures and let any type implement those methods, enabling flexible and extensible designs.
Implicit Implementation: Types do not need to declare that they implement an interface; it happens implicitly.
Method Sets: The interface defines a method set, and any type with methods matching that set satisfies the interface.
Dynamic Typing: Interface variables can hold values of different types that implement the interface.
*/

// 1- Defining Interfaces
//An interface is defined using the type keyword, followed by the interface name and the interface keyword.
/*
	type InterfaceName interface{
	  Method1(param1 Type1, param2 Type2) ReturnType
	  Method2(param1 Type1) ReturnType
	}
*/

//Example

// In this example, Shape is an interface that requires any implementing type to have Area and Perimeter methods.

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2- Implementing Interfaces
/*
To implement an interface, a type must implement all the methods declared in the interface.
There is no explicit declaration that a type implements an interface; it happens implicitly.
*/
//Example

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
	/////////Note:In Go, methods are functions that are associated with a particular type. They allow you to define behavior (methods) on user-defined types (like structs).
	/*
	   func (receiver ReceiverType) MethodName(parameters) returnType {
	       // method body
	   }
	*/

}

////////

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// 3- Interface Values
/*
1- Dynamic Type: The concrete type of the value assigned to the interface.
2- Dynamic Value: The actual value of the assigned variable.
*/
func interfaceValuesExample() {
	var s Shape
	s = Rectangle{Width: 5, Height: 4}
	fmt.Printf("Type: %T, Value: %v\n", s, s)
	// Output: Type: Rectangle, Value: {5 4}
	// Here, s is of type Shape, but at runtime, its dynamic type is Rectangle, and its dynamic value is Rectangle{5, 4}.
}

//4- Empty Interface (interface{})
//is an interface that has no methods. All types implement at least zero methods, so every type implements the empty interface.
/*
Note: While using interface{} provides flexibility, it requires type assertions or reflection to retrieve the underlying value,
which can introduce complexity and potential runtime errors.
*/

//5- Type Assertions
/*
A type assertion provides access to an interface's underlying concrete value.
value, ok := interfaceVariable.(ConcreteType)
value: The underlying value, if the assertion is successful.
ok: A boolean indicating whether the assertion succeeded.
*/
func typeInsertionExample() {
	var s Shape = Rectangle{Width: 5, Height: 4}
	rect, ok := s.(Rectangle)
	if ok {
		fmt.Println("Rectangle Width:", rect.Width)
	} else {
		fmt.Println("s is not a Rectangle")
	}
	//If the type assertion fails and you don't handle the ok value, it will cause a runtime panic.
}

// 6- Type Switches
//A type switch is a construct that allows you to perform several type assertions in series.

// Example
func describeShape(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Println("Shape is a Rectangle with width:", v.Width)
	case Circle:
		fmt.Println("Shape is a Circle with radius:", v.Radius)
	default:
		fmt.Println("Unknown shape")
	}
}

// 6- Interface Embedding
// Interfaces can embed other interfaces, combining their method sets.

type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
	//Here, ReadWriteCloser embeds io.Reader, io.Writer, and io.Closer, so any type that implements all methods from these interfaces satisfies ReadWriteCloser.
}

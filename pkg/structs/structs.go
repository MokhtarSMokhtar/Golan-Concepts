package structs

import "fmt"

//Struct in go
/*
In Go, a struct is a sequence of named elements called fields,
each field has a name and type. The name of a field must be unique within the struct.
Structs can be compared with classes in the Object-Oriented Programming paradigm.
type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}
*/
//Example

type Person struct {
	Name string
	Age  int
}

// 1- Field Name
/*
Field names in struct follow the Go convention
fields whose name starts with a lower case letter are only visible to code in the same package,
whereas those whose name starts with an upper case letter are visible in other packages.
*/

// 2- Create an instance from struct
/*
Field names in struct follow the Go convention
fields whose name starts with a lower case letter are only visible to code in the same package,
whereas those whose name starts with an upper case letter are visible in other packages.
*/
func createStructInstance() {
	// 1. Using the var Keyword
	var p Person
	p.Name = "Alice"
	p.Age = 30

	// 2. Using a Struct Literal
	p2 := Person{
		Name: "Bob",
		Age:  20,
	}
	fmt.Println(p2)

	// 3. Using the new Function
	P2 := new(Person) // new func return a pointer to the struct if ... for more a bout pointer https://go.dev/tour/moretypes/1
	p2.Name = "Charlie"
	p2.Age = 28
	fmt.Println(P2)

	// 4. Using an Anonymous Struct
	P3 := struct { // Anonymous structs are useful for quick, one-off data groupings without needing a named type.
		Name string
		Age  int
	}{
		Name: "Diana",
		Age:  22,
	}
	fmt.Println(P3)
}

//"New" functions
/*
Sometimes it can be nice to have functions that help us create struct instances.
By convention, these functions are usually called New or have their names starting with New,
but since they are just regular functions, you can give them any name you want.
They might remind you of constructors in other languages, but in Go they are just regular functions.
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
*/

//3- Accessing and Modifying Struct Fields
/*
You can access and modify struct fields using the dot (.) operator.
*/

func accessingStructFields() {
	// 1- via (.)
	p := Person{
		Name: "Eve",
		Age:  35,
	}

	fmt.Println(p.Name) // Outputs: Eve
	fmt.Println(p.Age)  // Outputs: 35

	// Modifying fields
	p.Age = 36
	fmt.Println(p.Age) // Outputs: 36

	// 2- Accessing via Pointer:
	//When you have a pointer to a struct,
	//you can access its fields using the dot operator without explicitly dereferencing the pointer.
	p2 := &Person{
		Name: "Frank",
		Age:  40,
	}

	fmt.Println(p2.Name) // Outputs: Frank
	p2.Age = 41
	fmt.Println(p2.Age) // Outputs: 41
}

// 3- Struct Methods

/*
 Go allows you to define methods associated with structs.
Methods enable you to define behaviors for your data types,
similar to how methods work in object-oriented programming.
func (receiver ReceiverType) MethodName(parameters) returnTypes {
    // Method body
}
*/

// Examples

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

/*
p := Person{
Name: "Grace",
Age: 29,
}
p.Greet() // Outputs: Hello, my name is Grace and I am 29 years old.

*/
//the next point should understand pointer and pass by value and pass by reference

// 4- Pointer vs. Value Receivers:
/*
Value Receiver (p Person): Receives a copy of the struct. Modifications within the method do not affect the original struct.
Pointer Receiver (p *Person): Receives a pointer to the struct. Modifications within the method can affect the original struct.

func (p *Person) HaveBirthday() {
    p.Age += 1
}

p := &Person{
    Name: "Henry",
    Age: 31,
}

p.HaveBirthday()
fmt.Println(p.Age) // Outputs: 32

*/
// -5 Struct Embedding (Composition)
/*
Go does not support traditional inheritance found in other object-oriented languages.
Instead, it uses composition through struct embedding to achieve similar functionality.
Struct Embedding:
Embedding one struct within another allows the outer struct to inherit the fields and methods of the embedded struct.
*/
// Example

type Employee struct {
	Person
	Position string
}

var emp = Employee{
	Person: Person{
		Name: "Ivy",
		Age:  27,
	},
	Position: "Developer",
}

//NOTE
/*
Overriding Methods:

If the outer struct defines a method with the same name as an embedded struct, it overrides the embedded method.
*/

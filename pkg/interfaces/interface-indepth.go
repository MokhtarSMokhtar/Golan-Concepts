package interfaces

import "fmt"

//////////Why We Need Interfaces in Go
// 1- Example Without Interfaces

//Consider the following types representing different animals:

type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

// Using the Types Separately:
func main() {
	c := Cat{}
	fmt.Println("Cat says:", c.Speak())

	d := Dog{}
	fmt.Println("Dog says:", d.Speak())
	//Cat says: meow
	//Dog says: woof
}

//=> While this works, if we wanted to handle both Cat and Dog uniformly
//(e.g., store them in a slice and iterate over them), we run into limitations because they are different types.

//Introducing Interfaces
//We can define an interface that captures the common behavior:

type Speaker interface {
	Speak() string
}

//Any type that has a Speak method with the same signature automatically satisfies the Speaker interface.
//Both Cat and Dog already have the required method, so they implement Speaker implicitly.

// Example after interface

func ExampleWithInterface() {
	animals := []Speaker{Cat{}, Dog{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	//meow
	//woof
}

/////////////Reducing Code Repetition
//Suppose we want to write a function that makes any animal speak:

//Without interface

func MakeCatSpeak(c Cat) {
	fmt.Println("Cat says:", c.Speak())
}

func MakeDogSpeak(d Dog) {
	fmt.Println("Dog says:", d.Speak())
}

//With Interfaces:

func MakeSpeak(s Speaker) {
	fmt.Println("Animal says:", s.Speak())
}

////////Why Interface Implementation is Implicit in Go
/*
Interface implementation is implicit in Go to promote simplicity and flexibility:
any type that implements the methods of an interface automatically satisfies it without explicit declaration.
This design reduces boilerplate code and decouples types from interfaces, enabling more modular and maintainable code.
*/

package maps

import "fmt"

//A map in Go is a built-in data type that associates keys with values.
//Maps are unordered collections, and the keys are unique within a map.
/*
Dynamic Size: Maps grow and shrink as needed.
Reference Type: Maps are reference types; when you pass them to functions or assign them to new variables, they reference the same underlying data.
Key Uniqueness: Each key in a map must be unique.
*/
//Note that slices and maps are exceptions to the above-mentioned rule.
//When we pass a slice or a map as arguments into a function, they are treated as pointer types even though there is no explicit * in the type.
//This means that if we pass a slice or map into a function and modify its underlying data, the changes will be reflected on the original slice or map.
/////////////////////////

// 1- Syntax
// map[KeyType]ValueType
// KeyType: The type of the keys (must be comparable).
// ValueType: The type of the values.
var personAge = map[string]int{
	"Alice": 30,
	"Bob":   25,
}

// 2- Declaring and Initializing Maps
func declaringAndInitialize() {
	// A- using make func
	//mapVariable := make(map[KeyType]ValueType)
	mapFromMake := make(map[string]int)
	fmt.Println(mapFromMake)
	// B- Using Map Literals
	//mapVariable := map[KeyType]ValueType{
	//    key1: value1,
	//    key2: value2,
	//    // ...
	//}
	personMap := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println(personMap)
}

// 3- Adding and update
func addingAndUpdate() {
	//You can add or update elements using the assignment syntax.
	//mapVariable[key] = value
	personAge["Alice"] = 44
}

// 4- Accessing Elements and iteration
func accessingElementAndIteration() {
	//value := mapVariable[key]
	age := personAge["Bob"] // Retrieves Bob's age
	fmt.Println(age)
	//checking for exist
	//value, exists := mapVariable[key]
	age, exists := personAge["Diana"]
	if exists {
		fmt.Println("Diana's age is", age)
	} else {
		fmt.Println("Diana is not in the map")
	}

	///Iterating Over Maps
	for name, age := range personAge {
		fmt.Printf("%s is %d years old\n", name, age)
	}

}

// 5- Deleting Elements
func deleteFromMap() {
	delete(personAge, "Diana")
}

//Maps as Reference Types
//Maps are reference types, similar to slices.
//Assigning a map to a new variable points both variables to the same underlying data.

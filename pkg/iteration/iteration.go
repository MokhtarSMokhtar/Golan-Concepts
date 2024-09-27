package iteration

import "fmt"

/*
 Looping is a fundamental concept in programming that allows you to execute a block of code repeatedly.
In Go, looping is primarily achieved
*/

// 1- For loop
/*
for initialization; condition; post {
    // Loop body
}
*/
func testForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)

	}
}

// 2. While Loop Style
// When you want to loop until a condition is false but don't need an initialization or post statement.
func whileLoopStyle() {
	i := 0
	for i < 5 {
		fmt.Println("Iteration:", i)
		i++
	}
}

// Range-Based Loops
// The range keyword allows you to iterate over elements in a collection
// like arrays, slices, maps, strings, and channels.
func testRangeLoop() {
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}

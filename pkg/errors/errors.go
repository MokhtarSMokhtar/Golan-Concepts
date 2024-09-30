package errors

import (
	"errors"
	"fmt"
	"os"
)

//Error handling in Go is designed to be simple and straightforward.
//Go uses error values to indicate an abnormal state. Instead of using exceptions,
//Go encourages explicit error checking and handling.

//NOTE :
//Errors are values of type error.
//Functions and methods that can fail return an error value along with any other results.
//The caller is responsible for checking the error and handling it appropriately.

//1- The Built-in error Interface
//In Go, the error type is an interface defined in the built-in errors package:
//type error interface {
//    Error() string
//}

//Any type that implements the Error() method with the signature Error() string satisfies the error interface.
//This means you can create custom error types by defining types that implement this method.
//Example:
//func (e MyError) Error() string {
//    return "an error occurred"
//}

// 2- Creating Errors
func createError() {
	//Using errors.New()
	err := errors.New("an error occurred")
	if err != nil {
		fmt.Println(err)

	}
	//Using fmt.Errorf()
	err1 := fmt.Errorf("an error occurred: %v", "something went wrong")
	if err1 != nil {
		fmt.Println(err)
	}
}

// 3- Custom Error Types

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
	return &MyError{Code: 404, Message: "Resource not found"}
}

func testCustomError() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}

// 4. Type Assertion for Errors
func openFile() error {
	_, err := os.Open("nonexistentfile.txt")
	return err
}
func testErrorTypeAssertion() {
	if err := openFile(); err != nil {
		// Checking for a specific error type
		if os.IsNotExist(err) {
			fmt.Println("File not found")
		} else {
			fmt.Println("Error:", err)
		}
	}
}

// Panic and Recover
// panic is a built-in function that stops the normal execution of a goroutine.
// It is used to indicate that something went unexpectedly wrong.
func checkValue(val int) {
	if val < 0 {
		panic("negative value")
	}
}

// Using recover
// recover regains control of a panicking goroutine. It can only be used inside deferred functions.
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	// Code that might panic
	panic("something bad happened")
}

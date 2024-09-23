package variable

/*
1- Variable Declaration
Go provides two primary ways to declare variables:
Using the "var" Keyword
Using Short Variable Declaration "(:=)"
*/

//////1- using var /////
/*
var variableName type
var variableName type = value
var variableName = value //
*/
//Examples
var age int
var name string = "Your Name"
var isStudent = true

//////1- using Short Variable Declaration /////
/*
variableName := value
Only usable inside functions.
Automatically infers the type based on the assigned value.
*/
//Examples
// score := 95.5 uncomment will cases an error

////////////////////////////////////////////////////////////////////////////////////////////////
/*
2- Type Inference
Go can automatically determine the type of a variable based on the value assigned to it
*/

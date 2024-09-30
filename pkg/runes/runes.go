package runes

import "fmt"

//Introduction to Runes
//In Go, a rune is an alias for int32 and is used to represent a Unicode code point.
//A Unicode code point is a numerical value that maps to a specific character in the Unicode standard,
//which encompasses characters from most of the world's writing systems.
//type rune = int32

//Some Points
//string type is a sequence of bytes, not necessarily characters.
//Runes provide a way to handle Unicode characters that may span multiple bytes.
//Using runes allows for proper manipulation and processing of Unicode strings.
//It is commonly used to store individual characters.

///Understanding Unicode and UTF-8 Encoding
//Unicode: A computing industry standard for consistent encoding, representation, and handling of text expressed in most of the world's writing systems.
//
//UTF-8: A variable-width character encoding capable of encoding all possible Unicode code points using one to four one-byte (8-bit) code units.
//Example:
//The character 'A' has a Unicode code point of U+0041 and is represented in UTF-8 as a single byte 0x41.

///Strings vs. Runes in Go

// Strings: In Go, strings are sequences of bytes ([]byte). They may contain text encoded in UTF-8.
func stringInGo() {
	s := "Hello"
	fmt.Println(len(s)) // Outputs the number of bytes
}

// Runes:Represent individual Unicode code points, allowing you to work with characters regardless of their byte length.
func runeInGo() {
	s := "Hello"
	r := []rune(s)
	fmt.Println(len(r)) // Outputs the number of runes (characters)

}

///Why is this important?
//Using len(s) on a string returns the number of bytes, not the number of characters.
//Multibyte characters (like many non-ASCII characters) can cause issues when processing strings byte by byte.

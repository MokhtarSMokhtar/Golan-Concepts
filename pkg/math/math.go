package main

import (
	"fmt"
	"math"
)

// Explain Go's math package with examples

// Basic Functions

// Absolute Value: Returns the absolute value of x.

func Abs() {
	x := -42.5
	fmt.Println("Absolute value:", math.Abs(x)) // Outputs: 42.5
}

// Rounding Functions

func RoundingFunctions() {
	// 1- math.Floor(x float64) float64
	// Returns the greatest integer value less than or equal to x.
	fmt.Println("Floor of 2.7:", math.Floor(2.7)) // Outputs: 2

	// 2- math.Ceil(x float64) float64
	// Returns the smallest integer value greater than or equal to x.
	fmt.Println("Ceil of 2.7:", math.Ceil(2.7)) // Outputs: 3

	// 3- math.Round(x float64) float64
	// Rounds x to the nearest integer, half away from zero.
	fmt.Println("Round of 2.5:", math.Round(2.5)) // Outputs: 3
	fmt.Println("Round of 2.4:", math.Round(2.4)) // Outputs: 2
}

// Maximum and Minimum
func MaxAndMin() {
	// math.Max(x, y float64) float64
	// Returns the larger of x or y.
	fmt.Println("Max of 10 and 20:", math.Max(10, 20)) // Outputs: 20

	// math.Min(x, y float64) float64
	// Returns the smaller of x or y.
	fmt.Println("Min of 10 and 20:", math.Min(10, 20)) // Outputs: 10
}

// Exponential and Logarithmic Functions
func ExponentialAndLogarithmicFunctions() {
	// math.Exp(x float64) float64
	// Returns eˣ, the base-e exponential of x.
	fmt.Println("Exp of 1:", math.Exp(1)) // Outputs: e ≈ 2.71828

	// math.Log(x float64) float64
	// Returns the natural logarithm (base e) of x.
	fmt.Println("Log of e:", math.Log(math.E)) // Outputs: 1

	// math.Log10(x float64) float64
	// Returns the decimal logarithm (base 10) of x.
	fmt.Println("Log10 of 1000:", math.Log10(1000)) // Outputs: 3
}

// Power Functions

func PowerFunctions() {
	// math.Pow(x, y float64) float64
	// Returns x raised to the power y (xʸ).
	fmt.Println("2 raised to the power 3:", math.Pow(2, 3)) // Outputs: 8

	// math.Sqrt(x float64) float64
	// Returns the square root of x.
	fmt.Println("Square root of 16:", math.Sqrt(16)) // Outputs: 4

	// math.Cbrt(x float64) float64
	// Returns the cube root of x.
	fmt.Println("Cube root of 27:", math.Cbrt(27)) // Outputs: 3
}

// Trigonometric Functions
func TrigonometricFunctions() {
	// math.Sin(x float64) float64
	// Returns the sine of x (x in radians).
	fmt.Println("Sin of Pi/2:", math.Sin(math.Pi/2)) // Outputs: 1

	// math.Cos(x float64) float64
	// Returns the cosine of x (x in radians).
	fmt.Println("Cos of Pi:", math.Cos(math.Pi)) // Outputs: -1

	// math.Tan(x float64) float64
	// Returns the tangent of x (x in radians).
	fmt.Println("Tan of Pi/4:", math.Tan(math.Pi/4)) // Outputs: 1
}

// Inverse Trigonometric Functions
func InverseTrigonometricFunctions() {
	// math.Asin(x float64) float64
	// Returns the arcsine of x in radians.
	fmt.Println("Asin of 0.5:", math.Asin(0.5)) // Outputs: Pi/6 ≈ 0.523598

	// math.Acos(x float64) float64
	// Returns the arccosine of x in radians.
	fmt.Println("Acos of 0.5:", math.Acos(0.5)) // Outputs: Pi/3 ≈ 1.047197

	// math.Atan(x float64) float64
	// Returns the arctangent of x in radians.
	fmt.Println("Atan of 1:", math.Atan(1)) // Outputs: Pi/4 ≈ 0.785398
}

// Hyperbolic Functions

func HyperbolicFunctions() {
	// math.Sinh(x float64) float64
	// Returns the hyperbolic sine of x.
	fmt.Println("Sinh of 0:", math.Sinh(0)) // Outputs: 0

	// math.Cosh(x float64) float64
	// Returns the hyperbolic cosine of x.
	fmt.Println("Cosh of 0:", math.Cosh(0)) // Outputs: 1

	// math.Tanh(x float64) float64
	// Returns the hyperbolic tangent of x.
	fmt.Println("Tanh of 0:", math.Tanh(0)) // Outputs: 0
}

// Special Functions

func SpecialFunctions() {
	// math.Mod(x, y float64) float64
	// Returns the floating-point remainder of x/y.
	fmt.Println("Mod of 10 and 3:", math.Mod(10, 3)) // Outputs: 1

	// math.Hypot(p, q float64) float64
	// Returns sqrt(p*p + q*q), the Euclidean distance.
	fmt.Println("Hypotenuse of sides 3 and 4:", math.Hypot(3, 4)) // Outputs: 5

	// math.Trunc(x float64) float64
	// Returns the integer value of x, truncated toward zero.
	fmt.Println("Truncate 3.7:", math.Trunc(3.7))   // Outputs: 3
	fmt.Println("Truncate -3.7:", math.Trunc(-3.7)) // Outputs: -3
}

// Working with Infinity and NaN

func InfinityAndNaN() {
	// math.Inf(sign int) float64
	// Returns positive or negative infinity.
	posInf := math.Inf(1)
	negInf := math.Inf(-1)
	fmt.Println("Positive Infinity:", posInf)
	fmt.Println("Negative Infinity:", negInf)

	// math.IsInf(f float64, sign int) bool
	// Checks if f is infinity.
	fmt.Println("Is posInf infinity?", math.IsInf(posInf, 0)) // Outputs: true

	// math.NaN() float64
	// Returns a NaN (Not a Number) value.
	nanValue := math.NaN()
	fmt.Println("NaN value:", nanValue)

	// math.IsNaN(f float64) bool
	// Checks if f is NaN.
	fmt.Println("Is nanValue NaN?", math.IsNaN(nanValue)) // Outputs: true
}

// Constants
func Constants() {
	// Math constants provided by the math package
	fmt.Println("Pi:", math.Pi)       // Outputs: 3.141592653589793
	fmt.Println("E:", math.E)         // Outputs: 2.718281828459045
	fmt.Println("Phi:", math.Phi)     // Outputs: 1.618033988749895
	fmt.Println("Sqrt2:", math.Sqrt2) // Outputs: 1.4142135623730951
}

// Main function to call all examples
func main() {
	fmt.Println("---- Absolute Value ----")
	Abs()

	fmt.Println("\n---- Rounding Functions ----")
	RoundingFunctions()

	fmt.Println("\n---- Maximum and Minimum ----")
	MaxAndMin()

	fmt.Println("\n---- Exponential and Logarithmic Functions ----")
	ExponentialAndLogarithmicFunctions()

	fmt.Println("\n---- Power Functions ----")
	PowerFunctions()

	fmt.Println("\n---- Trigonometric Functions ----")
	TrigonometricFunctions()

	fmt.Println("\n---- Inverse Trigonometric Functions ----")
	InverseTrigonometricFunctions()

	fmt.Println("\n---- Hyperbolic Functions ----")
	HyperbolicFunctions()

	fmt.Println("\n---- Special Functions ----")
	SpecialFunctions()

	fmt.Println("\n---- Working with Infinity and NaN ----")
	InfinityAndNaN()

	fmt.Println("\n---- Constants ----")
	Constants()
}

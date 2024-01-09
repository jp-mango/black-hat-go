package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// When consecutive parameters share a type, you can omit the type from all but the last
func add(x, y int) int {
	return x + y
}

// a function can return multiple parameters
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//! using libraries
	fmt.Println("\n-- Libraries --")
	// using fmt
	fmt.Println("Printing to the console uses fmt library: Hello, Justin")
	// using time
	fmt.Println("Current time using 'time' library:", time.Now())
	// using math/rand
	fmt.Println("Random number 0 - 10 using math/rand library:", rand.Intn(10))
	// using math
	fmt.Println("The square root of 7, using math library:", math.Sqrt(7))
	fmt.Println()

	//! formatted strings
	fmt.Println("-- Formatted Strings --")
	/*
		?	%v: The value in a default format.
		?	%T: A Go-syntax representation of the type of the value.
		?	%d: Base 10 integer formatting.
		?	%f: Floating-point number.
		?	%s: String formatting.
		?	%q: Quoted string.
		?	%p: Pointer address.
	*/
	myFloat := 123.456789
	fmt.Println("myFloat:", myFloat)
	fmt.Printf("Floating-point (2 decimal places): %.2f\n", myFloat)
	fmt.Println()

	//! functions
	fmt.Println("-- Functions --")
	// calling a function
	fmt.Println("Calling 'add' function with 2 and 5 will return", add(2, 5))
	// calling swap function
	a, b := swap("hello", "world") // the := symbol is used for short variable declarations. It allows you to declare and initialize variables in a single line without explicitly specifying their types.
	fmt.Println("Swap function produces:", a, b)
	fmt.Println()

	//! variable declaration
	fmt.Println("-- Variable Declaration --")
	// "var" statement declares a list of variables (without initializers)
	var c, python, java bool
	fmt.Println(c, python, java)
	// (with initializers) If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var i, j int = 1, 2
	fmt.Println(i, j)
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	name1, age1 := "Justin", 26
	fmt.Println(name1, age1)
	fmt.Println()

	//! basic variable types
	fmt.Println("-- Basic Variables --")
	//? Boolean Type
	var isBool bool // A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.

	//? Integer Types
	var numInt int       // An int is a signed integer of platform-dependent size.
	var numInt8 int8     // An int8 is a signed 8-bit integer.
	var numInt16 int16   // An int16 is a signed 16-bit integer.
	var numInt32 int32   // An int32 is a signed 32-bit integer.
	var numInt64 int64   // An int64 is a signed 64-bit integer.
	var numUint uint     // An uint is an unsigned integer of platform-dependent size.
	var numUint8 uint8   // An uint8 is an unsigned 8-bit integer (byte).
	var numUint16 uint16 // An uint16 is an unsigned 16-bit integer.
	var numUint32 uint32 // An uint32 is an unsigned 32-bit integer.
	var numUint64 uint64 // An uint64 is an unsigned 64-bit integer.
	//? Special integer types
	var numByte byte // Alias for uint8
	var numRune rune // Alias for int32 (represents a Unicode code point)
	//? Floating Point Types
	var numFloat32 float32 // A float32 is a 32-bit floating-point number.
	var numFloat64 float64 // A float64 is a 64-bit floating-point number.
	//? Complex Number Types
	var numComplex64 complex64   // A complex64 is a complex number with float32 real and imaginary parts.
	var numComplex128 complex128 // A complex128 is a complex number with float64 real and imaginary parts.
	//? String Type
	var str string // A string type represents a string of characters. Strings are immutable in Go.
	// Print all the zero values of the variables
	fmt.Println(isBool, numInt, numInt8, numInt16, numInt32, numInt64, numUint, numUint8, numUint16, numUint32, numUint64, numByte, numRune, numFloat32, numFloat64, numComplex64, numComplex128, str)
	fmt.Println()

	//! constants
	fmt.Println("-- Constants --")
	// constants are declared like variables but with the const keyword
	const Pi = math.Pi
	fmt.Println("pi = ", Pi)
	fmt.Println()

	//! loops
	fmt.Println("-- Loops --")
	//Go only uses a for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // adds all of the numbers [0:10)

	// an example of "for as a while" doing the same thing as above
	total := 0
	count := 0
	for count < 10 {
		total += count
		count++
	}
	fmt.Println(total)
	fmt.Println()

	//! Flow Control
	fmt.Println("-- Flow Control --")
	//? if/else
	test_number := 5

	if test_number%2 == 0 {
		fmt.Println(test_number, "is even")
	} else {
		fmt.Println(test_number, "is odd")
	}

	//? switch statement
	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//? defer - used to postpone the execution of a function or a statement. multiple stacked defers operate in "last-in-first-out"
	defer fmt.Println("\n~End Notes~") // this line will execute just before main returns
	fmt.Println()

	//! pointers
	fmt.Println("-- Pointers --")
	// Declare an integer variable
	originalValue := 100
	// Declare a pointer to an integer
	var pointerToValue *int
	// Assign the address of originalValue to pointerToValue
	pointerToValue = &originalValue
	// Print the original value, the address of the original value, and the pointer value
	fmt.Println("Original Value:", originalValue)
	fmt.Println("Address of Original Value:", &originalValue)
	fmt.Println("Pointer Value:", pointerToValue)
	// Change the value at the pointer
	*pointerToValue = 200
	// Print the changed value
	fmt.Println("Changed Value:", originalValue)
	fmt.Println()

	//! structs
	fmt.Println("-- Structs --")
	type Person struct {
		Name string
		Age  int
		City string
	}
	person1 := Person{"Justin", 26, "Atlanta"}
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("City:", person1.City)
	fmt.Println()

	//! Arrays
	fmt.Println("-- Arrays --")
	// Declare and initialize an array in a single line
	primes := [3]int{2, 3, 5}
	fmt.Println("Length of prime number list: ", cap(primes))
	fmt.Println("Prime numbers:", primes)
	// iterate over an array and print each element
	for i, num := range primes {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}

	//? slicing
	numList := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("First 5:", numList[:5])

	//? dynamically sized arrays using make (allocates a zeroed array and returns a slice that refers to that array)
	dynArray := make([]int, 5) // makes an array of zeros with a length of 5
	fmt.Println(dynArray)

}

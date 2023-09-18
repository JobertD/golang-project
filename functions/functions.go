package main

import (
	"fmt"
	"strings"
)

/*
	1. Function Declaration and Syntax

	Functions in Golang have similar structures with Java.
	Equivalent to "methods" in Java

	func function_name (Parameter-list) (Return type)
	- func: keyword in Go language (used to create functions)
	- function_name: name of the function
	- Parameter_list: name and type of function parameters
	- Return_type: optional

*/

// 1a.) Basic function invocation (without parameters)
func greet() string {
	return "Welcome to Golang!"
}

// 1b.) Declaring the same type of parameters (structure doesn't have to be
// 		ddInt(firstInt int, secondInt int))

func addInt(firstInt, secondInt int) int {
	return firstInt + secondInt
}

// 1c.) Difference between floating-point values
func subtFloat(firstFloat, secondFloat float32) float32 {
	return firstFloat - secondFloat
}

// 1d.) alternate method for return statements ("naked" returns)
func rectangleArea(length, width float64) (area float64) {
	area = length * width
	return
}

// 1e.) multiple return statements
func finalGrade(midtermGrade, tentativeFG float64) (float64, string) {
	averageGrade := (midtermGrade + tentativeFG) / 2
	var gradeLetter string
	switch {
	case averageGrade > 90:
		gradeLetter = "A"
	case averageGrade < 90 && averageGrade > 85:
		gradeLetter = "B"
	case averageGrade < 85 && averageGrade > 70:
		gradeLetter = "C"
	default:
		gradeLetter = "F"
	}

	return averageGrade, gradeLetter
}

/*
2. Variadic Functions
- has varying number of arguments of the same type
- allows you to work with different numbers of arguments without explicitly specifying their count
*/

// 2a.) Joining multiple string
func joinString(elements ...string) string {
	return strings.Join(elements, "-")
}

// 2b.) Adding multiple integers
func summationOfInt(digit ...int) int {
	total := 0
	for _, num := range digit {
		total += num
	}
	return total
}

/*
3. Anonymous Functions
- also known as "function literals"
- functions that don't have a name
- can be defined inline within other functions or expressions
*/

/*
4. Defer Statements (delayed invocations)
- ensure that a function call is performed later
- often used to clean up resources or for tasks that should run regardless of how a function exits

  ** Notice how the first statement in the main function will be executed last
*/

// Main function
func main() {

	defer fmt.Println("\nDeferred statement.")
	// 1a
	greet()

	// 1b
	firstIntDigit := 548
	secondIntDigit := 363

	fmt.Println("\nAddition of integers")
	fmt.Println(firstIntDigit, "+", secondIntDigit, "=", addInt(firstIntDigit, secondIntDigit))

	// 1c
	var firstFloatDigit float32 = 123.456
	var secondFloatDigit float32 = 789.10

	fmt.Println("\nSubtraction of floating-point numbers")
	fmt.Println(firstFloatDigit, "-", secondFloatDigit, "=", subtFloat(firstFloatDigit, secondFloatDigit))

	// 1d
	var length float64 = 12
	var width float64 = 13

	fmt.Println("\nDimensions of a rectangle")
	fmt.Println("length:", length, "\nwidth:", width, "\narea:", rectangleArea(length, width))

	// 1e
	var midGrade float64 = 94.5
	var finGrade float64 = 85.97
	avg, letter := finalGrade(midGrade, finGrade)

	fmt.Println("\nGrade Computation")
	fmt.Println("Midterm Grade:", midGrade,
		"\nTentative Final Grade:", finGrade,
		"\nFinal Grade and Grade Letter:", avg, letter)

	// 2a
	fmt.Println("\nString concatenation using variadic function")
	fmt.Println(joinString("Go", "Language", "is", "fun"))
	fmt.Println(joinString("H", "E", "L", "L", "O"))

	// 2b
	result1 := summationOfInt(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)
	result2 := summationOfInt(100, 200, 300, 400, 500)
	result3 := summationOfInt(-3, -4, -5, -6)

	fmt.Println("\nSummation of digits 1-25:", result1)
	fmt.Println("Summation of multiples of 100 (up to 500):", result2)
	fmt.Println("Summation of negative digits:", result3)

	// 3a. Basic Anonymous Function
	func() {
		fmt.Println("\nThis is a result of an anonymous function. A function within a (main)function!")
	}()

	// 3b. Assigning anonymous function to a variable
	multiply := func(a, b int) int {
		return a * b
	}
	product := multiply(10, 30)
	fmt.Printf("The product of %d and %d is %d\n", 10, 30, product)

	// 3c. Anonymous function with closure
	closureExample := 5

	// Anonymous function that captures the outer variable 'closureExample'
	increment := func() int {
		closureExample++
		return closureExample
	}

	fmt.Println("\nx:", closureExample)    // Output: closureExample: 5
	fmt.Println("Increment:", increment()) // Output: Increment: 1, resulting in 6
	fmt.Println("x:", closureExample)      // Output: closureExample: 6 (value of 'closureExample' has been modified)

	/*
		This is an example of a closure,
		where the function "remembers" the state of the variables it references,
		even after it goes out of scope.
	*/

}

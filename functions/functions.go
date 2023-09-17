package main

import (
	"fmt"
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

// Basic function invocation (without parameters)
func greet() string {
	return "Welcome to Golang!"
}

// Declaring the same type of parameters (structure doesn't have to be
// addInt(firstInt int, secondInt int))

func addInt(firstInt, secondInt int) int {
	return firstInt + secondInt
}

func subtFloat(firstFloat, secondFloat float32) float32 {
	return firstFloat - secondFloat
}

// alternate method for return statements ("naked" returns)
func rectangleArea(length, width float64) (area float64) {
	area = length * width
	return
}

// multiple return statements
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
	2. Function Declaration and Syntax

	Functions in Golang have similar structures with Java.
	Equivalent to "methods" in Java

	func function_name (Parameter-list) (Return type)
	- func: keyword in Go language (used to create functions)
	- function_name: name of the function
	- Parameter_list: name and type of function parameters
	- Return_type: optional

*/

// Main function
func main() {

	greet()

	firstIntDigit := 548
	secondIntDigit := 363

	fmt.Println("\nAddition of integers")
	fmt.Println(firstIntDigit, "+", secondIntDigit, "=", addInt(firstIntDigit, secondIntDigit))

	var firstFloatDigit float32 = 123.456
	var secondFloatDigit float32 = 789.10

	fmt.Println("\nSubtraction of floating-point numbers")
	fmt.Println(firstFloatDigit, "-", secondFloatDigit, "=", subtFloat(firstFloatDigit, secondFloatDigit))

	var length float64 = 12
	var width float64 = 13
	fmt.Println("\nDimensions of a rectangle")
	fmt.Println("length:", length, "\nwidth:", width, "\narea:", rectangleArea(length, width))

	var midGrade float64 = 94.5
	var finGrade float64 = 85.97
	avg, letter := finalGrade(midGrade, finGrade)

	fmt.Println("\nGrade Computation")
	fmt.Println("Midterm Grade:", midGrade,
		"\nTentative Final Grade:", finGrade,
		"\nFinal Grade and Grade Letter:", avg, letter)

}

package controlflow

import (
	"fmt"
	"math"
)

func Controlflow() {

	// FOR LOOPS

	/*
		A for loop consists of three statements separated by semicolons and the looping body.
		The 3 statements consist of the following:
		1. The init statement - used to initialize variables whose scopes are within the for loop only
		2. The condition expression - used to check if the for loop should continue just before every iteration
		3. The post statement - executes after every iteration; used to increment or decrement variable

		Unlike in Java, these statements do not need to be enclosed in parentheses, but the body of the loop must always be enclosed in
		curly braces. Here's an example of a for loop in Go syntax
	*/
	fmt.Println("Running first example of for loop:")
	firstSum := 0
	for i := 0; i <= 10; i++ {
		firstSum += i
	}
	fmt.Println(firstSum)

	/*
		The statements used before the body of the for loop can be optional; thus you can omit one or more of the given expressions.
		If the 1st and 3rds statements are omitted, the semicolons do not need to be included. This effectively makes the for loop
		similar to a while loop. Additionally, if the 2nd statement is omitted, the conditional expression always evaluates to true and
		the loop will continue on unless the break statement is used.

	*/
	fmt.Println("Running second example of for loop:")
	secondSum, j := 0, 0
	for j <= 10 {
		secondSum += j
		j++
	}

	fmt.Println(secondSum)

	// IF-ELSE STATEMENTS

	// If statements in Go are similar to those in Java or C, except you do not need to enclose the conditional expression in parentheses;
	// the curly braces however are also important.

	fmt.Println("Running first example of if-else statements")
	x := float64(-4)
	var sqrtValue string
	if x < 0 {
		sqrtValue = fmt.Sprint(math.Sqrt(-x)) + "i"
	} else {
		sqrtValue = fmt.Sprint(math.Sqrt(x))
	}

	fmt.Printf("Square root of x = -4 is: %s\n", sqrtValue)

	// You may also make use of a short statement before the conditional which can be used to prepare variables within the scope of the
	// if-else blocks
	fmt.Println("Running second example of if-else statements")
	if power := math.Pow(2, 8); power <= 25 {
		fmt.Printf("%.2f is <= 25\n", power)
	} else if power <= 50 {
		fmt.Printf("%.2f is <= 50\n", power)
	} else if power <= 100 {
		fmt.Printf("%.2f is <= 100\n", power)
	} else {
		fmt.Printf("%.2f is > 100\n", power)
	}

	// SWITCH STATEMENTS
	/*
		Go's switch statements are similar to ones in Java or C, but each case does not have to be a constant representing an integer value.
		The switch cases may also not be followed by a break statement since only one case will always be executed.
	*/
	fmt.Println("Running example of switch statements")
	switch day := "monday"; day {
	case "saturday":
		fmt.Println("Today is Saturday!")
	case "monday":
		fmt.Println("Today is Monday!")
	default:
		fmt.Println("Today is " + day)
	}

}

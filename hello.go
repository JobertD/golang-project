package main

import "fmt"

func main() {
	menu()
}

func menu() {
	//can add more options if we want to show more examples/adjust if u add something
	var input int8
	//TODO make problems
	//FIXME break out of loop when one of switch case is triggered and proceed to method in switch case

	fmt.Println("1. ")
	fmt.Println("2. ")
	fmt.Println("3. ")
	fmt.Println("4. ")
	fmt.Println("5. ")
	fmt.Println("6. ")
	fmt.Println("7. ")
	fmt.Println("8. ")
	fmt.Println("9. ")
	fmt.Println("10. ")
	fmt.Println("11. Exit")
	fmt.Println("Please input a number: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		firstProblem()
	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 9:

	case 10:

	case 11:
		fmt.Println("Thank you for using our program...")
	default:
		fmt.Println("Please enter a valid number from the options menu.")
	}
}

func firstProblem() {
	var sagot string
	fmt.Println("Test?")
	fmt.Scanln(&sagot)

	fmt.Scanln()
	menu()
}

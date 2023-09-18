package main

import (
	"fmt"
	"presentation/compositeTypes"
	"presentation/controlflow"
)

func Main() {
	Menu()
}

func Menu() {
	//can add more options if we want to show more examples/adjust if u add something
	var input int8

	fmt.Println("1. Control Flow Statements")
	fmt.Println("2. Composite Types")
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
		fmt.Println()
		controlflow.Controlflow()
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		fmt.Scanln()
	case 2:
		fmt.Println()
		compositeTypes.Menu()
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		fmt.Scanln()
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
		break
	default:
		fmt.Println("Please enter a valid number from the options menu.")
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	menu()
}

func menu() {
	//can add more options if we want to show more examples/adjust if u add something
	var input int8

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
		firstExample()
	case 2:
		Menu()
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

func firstExample() {

	fmt.Println("Running first example of for loop:")
	firstSum := 0
	for i := 0; i <= 10; i++ {
		firstSum += i
	}
	fmt.Println(firstSum)

	fmt.Println("Running second example of for loop:")
	secondSum, j := 0, 0
	for j <= 10 {
		secondSum += j
		j++
	}

	fmt.Println(secondSum)

	fmt.Println("Running first example of if-else statements")
	x := float64(-4)
	var sqrtValue string
	if x < 0 {
		sqrtValue = fmt.Sprint(math.Sqrt(-x)) + "i"
	} else {
		sqrtValue = fmt.Sprint(math.Sqrt(x))
	}

	fmt.Printf("Square root of x = -4 is: %s\n", sqrtValue)

	if power := math.Pow(2, 8); power <= 25 {
		fmt.Printf("%.2f is <= 25\n", power)
	} else if power <= 50 {
		fmt.Printf("%.2f is <= 50\n", power)
	} else if power <= 100 {
		fmt.Printf("%.2f is <= 100\n", power)
	} else {
		fmt.Printf("%.2f is > 100\n", power)
	}

	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	fmt.Scanln()

	menu()
}

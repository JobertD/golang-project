package compositeTypes

import (
	"fmt"
)

func Menu() {
	// Display the menu options.
	fmt.Println("Menu:")
	fmt.Println("1. Arrays")
	fmt.Println("2. Slices")
	fmt.Println("3. Structs")
	fmt.Println("4. Maps")
	fmt.Println("5. Exit")

	var choice int
	// Read the user's choice.
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	// Use a switch-case to perform actions based on the choice.
	switch choice {
	case 1:
		array()
	case 2:
		slice()
	case 3:
		structs()
	case 4:
		maps()
	case 5:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
	Menu()
}

func array() {
	// Arrays

	var stringArray = [5]string{"Golang", "Mascot", "Is The"} // Initialize Array with 5 capacity
	fmt.Println("String Array: ", stringArray)
	stringArray[3] = "Best"               // Replace value on specified index
	stringArray[4] = "Gopher"             // Replace value on specified index
	intArray := [...]int{1, 2, 3, 10: 11} // Initialize Array with 11 capacity

	fmt.Println("String Array After Replace: ", stringArray)
	fmt.Println("Integer Array: ", intArray)
	fmt.Println("Integer Length: ", len(intArray))
}

func slice() {
	// Slices
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Integer Slice:", intSlice)
	intSlice = append(intSlice, 6)
	fmt.Println("Appended Integer Slice:", intSlice)

	part := intSlice[2:4]
	fmt.Println("Slice Length: ", len(part))
	fmt.Println("Slice Capacity: ", cap(part))
}

func structs() {
	// Structs

	// Create a struct
	type Person struct {
		Fname string
		Lname string
		Age   int
	}
	//Create an Object
	person1 := Person{
		Fname: "Himeko",
		Lname: "Murata",
		Age:   27,
	}
	//Create an Object
	person2 := Person{
		Fname: "Welt",
		Lname: "Yang",
		Age:   82,
	}
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)
}

func maps() {
	// Maps

	// Initialize a Map
	countryPopulation := map[string]int{
		"USA":     331002651,
		"India":   1380004385,
		"China":   1444216107,
		"Brazil":  212559417,
		"Nigeria": 206139587,
	}

	fmt.Println("USA Population:", countryPopulation["USA"])
	fmt.Println("India Population:", countryPopulation["India"])
	fmt.Println("China Population:", countryPopulation["China"])
	fmt.Println("Brazil Population:", countryPopulation["Brazil"])
	fmt.Println("Nigeria Population:", countryPopulation["Nigeria"])
}

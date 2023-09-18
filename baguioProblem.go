package main

import (
	"fmt"
)

const (
	totalParkingSlots = 10
)

type ParkingLot struct {
	Slots [totalParkingSlots]string
}

func main() {
	for {
		fmt.Println("Parking Reservation System")
		fmt.Println("1. Reserve Parking Slot")
		fmt.Println("2. Vacate Parking Slot")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			{
				displayLocations()
				var locationChoice int
				fmt.Print("Enter your location choice: ")
				fmt.Scanln(&locationChoice)

				switch locationChoice {
				case 1:
					// User chose Sm Baguio City
					displaySMParkingStations()
					var stationChoice int
					fmt.Print("Enter your station choice: ")
					fmt.Scanln(&stationChoice)

					switch stationChoice {
					case 1:
						// User chose Old Parking Area
						// Implement the reservation logic for Old Parking Area here
						fmt.Println("Reserved a parking slot in Sm Baguio City's Old Parking Area")
					case 2:
						// User chose New Parking Area
						// Implement the reservation logic for New Parking Area here
						fmt.Println("Reserved a parking slot in Sm Baguio City's New Parking Area")
					case 3:
						// User chose Outside Parking Area
						// Implement the reservation logic for Outside Parking Area here
						fmt.Println("Reserved a parking slot in Sm Baguio City's Outside Parking Area")
					default:
						fmt.Println("Invalid station choice. Please select a valid option.")
					}
				// Add more cases for other landmarks if needed
				default:
					fmt.Println("Invalid location choice. Please select a valid option.")
				}
			}
		case 2:
			{
				displayLocations()
			}
		case 3:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func displayLocations() {
	fmt.Println("Select Landmark")
	fmt.Println("1. Sm Baguio City")
	fmt.Println("2. Burnham Park")
	fmt.Println("3. Camp John Hay")
}

func displaySMParkingStations() {
	fmt.Println("Sm Baguio City")
	fmt.Println("1. Old Parking Area")
	fmt.Println("2. New Parking Area")
	fmt.Println("3. Outisde parking Area")
}

func generateRandomID() string {

}

func reserveParkingSlot(parkingLot *ParkingLot, id string) int {

}

func vacateParkingSlot(parkingLot *ParkingLot, slotNumber int) bool {

}

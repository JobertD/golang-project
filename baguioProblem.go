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

func generateRandomID() string {

}

func reserveParkingSlot(parkingLot *ParkingLot, id string) int {

}

func vacateParkingSlot(parkingLot *ParkingLot, slotNumber int) bool {

}

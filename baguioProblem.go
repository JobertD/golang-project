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
			}
		case 2:
			{
			}
		case 3:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func generateRandomID() string {

}

func reserveParkingSlot(parkingLot *ParkingLot, id string) int {

}

func vacateParkingSlot(parkingLot *ParkingLot, slotNumber int) bool {

}

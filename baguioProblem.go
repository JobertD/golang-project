package main

import (
	"fmt"
	"math/rand"
)

const (
	totalParkingSlots = 10
)

type ParkingLot struct {
	Slots [totalParkingSlots]string
}

func main() {

	parkingLot := &ParkingLot{}

	for {

		for i, slot := range parkingLot.Slots {
			if slot == "" {
				fmt.Printf("Slot %d: Available\n", i+1)
			} else {
				fmt.Printf("Slot %d: Taken by car with ID: %s\n", i+1, slot)
			}
		}
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
				var slotNumber int
				fmt.Print("Enter the slot number to reserve: ")
				fmt.Scanln(&slotNumber)
				id := generateRandomID()
				if reserveSpecificParkingSlot(parkingLot, slotNumber, id) {
					fmt.Printf("Reserved slot %d for car with ID: %s\n", slotNumber, id)
				} else {
					fmt.Printf("Slot %d is already taken or invalid.\n", slotNumber)

				}
				fmt.Println("Press any button to continue...")
				fmt.Scanln()
				fmt.Scanln()
			}
		case 2:
			{
				displayLocations()
				var slotNumber int
				fmt.Print("Enter the slot number to vacate: ")
				fmt.Scanln(&slotNumber)
				if vacateParkingSlot(parkingLot, slotNumber) {
					fmt.Printf("Slot %d is now vacant.\n", slotNumber)
				} else {
					fmt.Printf("Slot %d is already vacant or invalid.\n", slotNumber)

				}

				fmt.Println()
				fmt.Println("Press any button to continue...")
				fmt.Scanln()
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
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func reserveSpecificParkingSlot(parkingLot *ParkingLot, slotNumber int, id string) bool {
	if slotNumber >= 1 && slotNumber <= totalParkingSlots && parkingLot.Slots[slotNumber-1] == "" {
		parkingLot.Slots[slotNumber-1] = id
		return true
	}
	return false
}

func vacateParkingSlot(parkingLot *ParkingLot, slotNumber int) bool {
	if slotNumber >= 1 && slotNumber <= totalParkingSlots && parkingLot.Slots[slotNumber-1] != "" {
		parkingLot.Slots[slotNumber-1] = ""
		return true
	}
	return false
}

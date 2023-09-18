package main

import (
	"fmt"
	"math/rand"
)

// Create Location struct
type Location struct {
	Name           string
	TotalSlots     int
	AvailableSlots int
	ParkingLot     ParkingLot
}

// Create ParkingLot struct
type ParkingLot struct {
	Slots []string
}

// Initialize Variables
var (
	TotalSlotsSmBaguio    = 20
	TotalSlotsBurnhamPark = 15
	TotalSlotsCampJohnHay = 10
)

// Main Method
func main() {
	BaguioMenu()
}

func BaguioMenu() {
	// Object Creations
	smBaguio := Location{
		Name:           "Sm Baguio City",
		TotalSlots:     TotalSlotsSmBaguio,
		AvailableSlots: TotalSlotsSmBaguio,
		ParkingLot:     ParkingLot{Slots: make([]string, TotalSlotsSmBaguio)},
	}

	burnhamPark := Location{
		Name:           "Burnham Park",
		TotalSlots:     TotalSlotsBurnhamPark,
		AvailableSlots: TotalSlotsBurnhamPark,
		ParkingLot:     ParkingLot{Slots: make([]string, TotalSlotsBurnhamPark)},
	}

	campJohnHay := Location{
		Name:           "Camp John Hay",
		TotalSlots:     TotalSlotsCampJohnHay,
		AvailableSlots: TotalSlotsCampJohnHay,
		ParkingLot:     ParkingLot{Slots: make([]string, TotalSlotsCampJohnHay)},
	}

	// Creates an array
	locations := []Location{smBaguio, burnhamPark, campJohnHay}

	for {
		displayLocations(locations) // Display Locations
		var locationChoice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&locationChoice) // Scans Input
		if locationChoice < 1 || locationChoice > len(locations) {
			fmt.Println("Invalid location choice. Please select a valid option.")
			continue
		}

		// Picks the Location from array
		location := &locations[locationChoice-1]

		for {
			fmt.Printf("Parking Slots for %s:\n", location.Name)

			// Prints out the list of the slots for the chosen location
			for i, slot := range location.ParkingLot.Slots {
				if slot == "" {
					fmt.Printf("Slot %d: Available\n", i+1)
				} else {
					fmt.Printf("Slot %d: Taken by car with ID: %s\n", i+1, slot)
				}
			}

			// Prints Menu
			fmt.Println("Parking Reservation System")
			fmt.Println("1. Reserve Parking Slot")
			fmt.Println("2. Vacate Parking Slot")
			fmt.Println("3. Go back")

			var choice int
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice) // Scans Input

			switch choice {
			case 1:
				var slotNumber int
				fmt.Print("Enter the slot number to reserve: ")
				fmt.Scanln(&slotNumber)  // Scans Input
				id := generateRandomID() // Generate Random ID
				if reserveSpecificParkingSlot(location, slotNumber, id) {
					fmt.Printf("Reserved slot %d for car with ID: %s\n", slotNumber, id)
				} else {
					fmt.Printf("Slot %d is already taken or invalid.\n", slotNumber)
				}

			case 2:
				var slotNumber int
				fmt.Print("Enter the slot number to vacate: ")
				fmt.Scanln(&slotNumber) // Scans Input
				if vacateParkingSlot(location, slotNumber) {
					fmt.Printf("Slot %d is now vacant.\n", slotNumber)
				} else {
					fmt.Printf("Slot %d is already vacant or invalid.\n", slotNumber)
				}
			case 3:
				fmt.Println("Going back to the main menu.")
				BaguioMenu() // Return to menu
			default:
				fmt.Println("Invalid choice. Please select a valid option.")
			}
		}
	}
}

// Displays the available locations
func displayLocations(locations []Location) {
	fmt.Println("Select Location:")
	for i, location := range locations {
		fmt.Printf("%d. %s \n", i+1, location.Name)
	}
}

// Generate Random ID
func generateRandomID() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Reserve A Slot
func reserveSpecificParkingSlot(location *Location, slotNumber int, id string) bool {
	if slotNumber >= 1 && slotNumber <= location.TotalSlots && location.ParkingLot.Slots[slotNumber-1] == "" {
		location.ParkingLot.Slots[slotNumber-1] = id
		location.AvailableSlots--
		return true
	}
	return false
}

// Vacate Slot
func vacateParkingSlot(location *Location, slotNumber int) bool {
	if slotNumber >= 1 && slotNumber <= location.TotalSlots && location.ParkingLot.Slots[slotNumber-1] != "" {
		location.ParkingLot.Slots[slotNumber-1] = ""
		location.AvailableSlots++
		return true
	}
	return false
}

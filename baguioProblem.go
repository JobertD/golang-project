package main

import (
	"fmt"
	"math/rand"
)

type Location struct {
	Name           string
	TotalSlots     int
	AvailableSlots int
	ParkingLot     ParkingLot
}

const (
	TotalSlotsSmBaguio    = 20
	TotalSlotsBurnhamPark = 15
	TotalSlotsCampJohnHay = 10
)

type ParkingLot struct {
	Slots []string
}

func main() {
	BaguioMenu()
}

func BaguioMenu() {
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

	locations := []Location{smBaguio, burnhamPark, campJohnHay}

	for {
		displayLocations(locations)
		var locationChoice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&locationChoice)
		if locationChoice < 1 || locationChoice > len(locations) {
			fmt.Println("Invalid location choice. Please select a valid option.")
			continue
		}

		location := &locations[locationChoice-1]

		for {
			fmt.Printf("Parking Slots for %s:\n", location.Name)
			for i, slot := range location.ParkingLot.Slots {
				if slot == "" {
					fmt.Printf("Slot %d: Available\n", i+1)
				} else {
					fmt.Printf("Slot %d: Taken by car with ID: %s\n", i+1, slot)
				}
			}

			fmt.Println("Parking Reservation System")
			fmt.Println("1. Reserve Parking Slot")
			fmt.Println("2. Vacate Parking Slot")
			fmt.Println("3. Go back")

			var choice int
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				var slotNumber int
				fmt.Print("Enter the slot number to reserve: ")
				fmt.Scanln(&slotNumber)
				id := generateRandomID()
				if reserveSpecificParkingSlot(location, slotNumber, id) {
					fmt.Printf("Reserved slot %d for car with ID: %s\n", slotNumber, id)
					location.AvailableSlots -= 1
				} else {
					fmt.Printf("Slot %d is already taken or invalid.\n", slotNumber)
				}
			case 2:
				var slotNumber int
				fmt.Print("Enter the slot number to vacate: ")
				fmt.Scanln(&slotNumber)
				if vacateParkingSlot(location, slotNumber) {
					fmt.Printf("Slot %d is now vacant.\n", slotNumber)
					location.AvailableSlots += 1
				} else {
					fmt.Printf("Slot %d is already vacant or invalid.\n", slotNumber)
				}
			case 3:
				fmt.Println("Going back to the main menu.")
				BaguioMenu()
			default:
				fmt.Println("Invalid choice. Please select a valid option.")
			}
		}
	}
}

func displayLocations(locations []Location) {
	fmt.Println("Select Location:")
	for i, location := range locations {
		fmt.Printf("%d. %s (%d available slots)\n", i+1, location.Name, location.AvailableSlots)
	}
}

func generateRandomID() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func reserveSpecificParkingSlot(location *Location, slotNumber int, id string) bool {
	if slotNumber >= 1 && slotNumber <= location.TotalSlots && location.ParkingLot.Slots[slotNumber-1] == "" {
		location.ParkingLot.Slots[slotNumber-1] = id
		location.AvailableSlots--
		return true
	}
	return false
}

func vacateParkingSlot(location *Location, slotNumber int) bool {
	if slotNumber >= 1 && slotNumber <= location.TotalSlots && location.ParkingLot.Slots[slotNumber-1] != "" {
		location.ParkingLot.Slots[slotNumber-1] = ""
		location.AvailableSlots++
		return true
	}
	return false
}

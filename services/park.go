package services

import (
	"fmt"
	"strconv"
	"vimaldumdum/parklot/models"
	"vimaldumdum/parklot/utility"
)

func NewParkService(command []string, lot *models.ParkingLot) {
	// fmt.Println("inside park service")
	switch command[0] {
	case "park_vehicle":
		park(command[1], command[2], command[3], lot)
		break
	case "unpark_vehicle":
		unpark(command[1], lot)
	}
}

func park(vehicleType, reg, color string, lot *models.ParkingLot) {
	// fmt.Println("inside park func")
	lotId, floor, slot, vType, storage, floors, slots := lot.GetId(), -1, -1, utility.GetVehicleType(vehicleType), lot.GetStorage(), lot.GetFloorNumber(), lot.GetSlotNumber()
	totalSlots := utility.GetAllSlotsForVehicle(vType, slots)
	start, _ := utility.GetStartEndForVehicleType(vType, slots)
	for i := 0; i < floors; i++ {
		if storage[i][vType] == totalSlots {
			continue
		}
		floor = i + 1
		slot = storage[i][vType] + start
		storage[i][vType]++
		break
	}
	if floor == -1 || slot == -1 {
		fmt.Println("Parking Lot Full")
	} else {
		strFloor := fmt.Sprintf("%d", floor)
		strSlot := fmt.Sprintf("%d", slot)
		fmt.Printf("Parked vehicle. Ticket ID: %s\n", utility.GenerateTicket(lotId, strFloor, strSlot))
	}
}

func unpark(ticket string, lot *models.ParkingLot) {
	f, s := utility.GetFloorAndSlotFromTicket(ticket)
	floor, _ := strconv.Atoi(f)
	slot, _ := strconv.Atoi(s)
	v := -1
	//get vehicle from slot number
	if slot < 2 {
		v = 0
	} else if slot < 4 {
		v = 1
	} else {
		v = 2
	}
	storage := lot.GetStorage()

	storage[floor-1][v]--

	fmt.Println("Unparked vehicle")
}

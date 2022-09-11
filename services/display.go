package services

import (
	"fmt"
	"vimaldumdum/parklot/models"
	"vimaldumdum/parklot/utility"
)

func NewDisplayService(displayType, vehicleType string, lot *models.ParkingLot) {

	storage := lot.GetStorage()
	floors := lot.GetFloorNumber()
	slots := lot.GetSlotNumber()
	vType := utility.GetVehicleType(vehicleType)
	totalSlots := utility.GetAllSlotsForVehicle(vType, slots)

	// fmt.Println("inside display service")

	switch displayType {
	case "free_count":
		freeCount(storage, floors, totalSlots, vType, vehicleType)
		break
	case "free_slots":
		freeSlots(storage, floors, vType, totalSlots, vehicleType)
		break
	case "occupied_slots":
		occupiedSlots(storage, floors, vType, totalSlots, vehicleType)
		break
	}

}

func occupiedSlots(storage [][3]int, floors, vType, totalSlots int, vehicleType string) {
	start, _ := utility.GetStartEndForVehicleType(vType, totalSlots)

	for i := 0; i < floors; i++ {
		fmt.Printf("Occupied slots for %s on Floor %d: ", vehicleType, i+1)
		for j := start; j < start+storage[i][vType]; j++ {
			fmt.Printf("%d, ", j)
		}
		fmt.Println()
	}
}

func freeSlots(storage [][3]int, floors, vType, totalSlots int, vehicleType string) {
	start, end := utility.GetStartEndForVehicleType(vType, totalSlots)

	for i := 0; i < floors; i++ {
		fmt.Printf("Free slots for %s on Floor %d: ", vehicleType, i+1)
		for j := start + storage[i][vType]; j <= end; j++ {
			fmt.Printf("%d, ", j)
		}
		fmt.Println()
	}
}

func freeCount(storage [][3]int, floors, totalSlots, vType int, vehicleType string) {
	// fmt.Println("inside free count display service")
	for i := 0; i < floors; i++ {
		count := totalSlots - storage[i][vType]
		fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, i+1, count)
	}
}

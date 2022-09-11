package models

type ParkingLot struct {
	id             string
	numberOfFloors int
	numberOfSlots  int
	floor          [][3]int
}

func NewParkingLot(floors, slots int, id string) *ParkingLot {
	return &ParkingLot{
		id:             id,
		numberOfFloors: floors,
		numberOfSlots:  slots,
		floor:          make([][3]int, floors),
	}
}

func (lot *ParkingLot) GetStorage() [][3]int {
	return lot.floor
}

func (lot *ParkingLot) GetFloorNumber() int {
	return lot.numberOfFloors
}

func (lot *ParkingLot) GetSlotNumber() int {
	return lot.numberOfSlots
}

func (lot *ParkingLot) GetId() string {
	return lot.id
}

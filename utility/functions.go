package utility

import (
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader

func GetReader() *bufio.Reader {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	return reader
}

func GetAllSlotsForVehicle(v int, slots int) int {

	switch v {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return slots - 3
	}

	return 0
}

func GetVehicleType(t string) int {

	switch t {
	case "TRUCK":
		return 0
	case "BIKE":
		return 1
	case "CAR":
		return 2
	}
	return -1
}

func GetStartEndForVehicleType(v, slots int) (int, int) {

	switch v {
	case 0:
		return 1, 1
	case 1:
		return 2, 3
	case 2:
		return 4, slots
	}

	return -1, -1
}

func GenerateTicket(lotId, floor, slot string) string {
	return lotId + "_" + floor + "_" + slot
}

func GetFloorAndSlotFromTicket(ticket string) (string, string) {
	words := strings.Split(ticket, "_")
	return words[1], words[2]
}

package services

import (
	"strings"

	"vimaldumdum/parklot/models"
	"vimaldumdum/parklot/utility"
)

func NewParkingLotService(id string, floors, slots int) {
	lot := models.NewParkingLot(floors, slots, id)

	var command string

	for true {
		command, _ = utility.GetReader().ReadString('\n')
		command = strings.TrimSuffix(command, "\n")
		// fmt.Printf("Command enterd: %s\n", command)
		if command == "exit" {
			break
		}
		words := strings.Split(command, " ")

		// fmt.Printf("%s,%s,%s\n", words[0], words[1], words[2])

		// fmt.Println(len(wexitords))
		if words[0] == "display" {
			// fmt.Println("inside display block")
			NewDisplayService(words[1], words[2], lot)
		} else {
			NewParkService(words, lot)
		}
	}
}

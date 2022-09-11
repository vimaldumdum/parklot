package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"vimaldumdum/parklot/services"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {

	defer handlePanic()
	fmt.Printf("Create parking lot: \n")
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSuffix(inp, "\n")
	command := strings.Split(inp, " ")

	floors, _ := strconv.Atoi(command[2])
	slots, _ := strconv.Atoi(command[3])
	services.NewParkingLotService(command[1], floors, slots)
}

func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER: ", a)
		debug.PrintStack()
	}
}

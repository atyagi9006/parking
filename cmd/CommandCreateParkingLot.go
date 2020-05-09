package cmd

import (
	"fmt"
	"strconv"
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandCreateParkingLot struct {
	Command
	Capacity uint
}

func (cmd *CommandCreateParkingLot) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	value, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		return errInvalidParam
	}
	cmd.Capacity = uint(value)
	return nil
}

func (cmd *CommandCreateParkingLot) Clear() {
	cmd.Args = nil
	cmd.Capacity = 0
}

func (cmd *CommandCreateParkingLot) ValidateParams() bool {
	return len(cmd.Args) == 1
}

func (cmd *CommandCreateParkingLot) Run() (string, error) {
	var output string
	obj := parking.New(cmd.Capacity)
	if obj == nil {
		return output, fmt.Errorf("Error")
	}
	output = fmt.Sprintf("Created a parking lot with %d slots", cmd.Capacity)
	return output, nil
}

package cmd

import (
	"fmt"
	"strings"

	"github.com/atyagi9006/parking/car"
	parking "github.com/atyagi9006/parking/lot"
)

type CommandPark struct {
	Command
	Car *car.Car
}

func (cmd *CommandPark) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	cmd.Car = car.New(cmd.Args[0], cmd.Args[1])
	return nil
}

func (cmd *CommandPark) Clear() {
	cmd.Args = nil
	cmd.Car = nil
}

func (cmd *CommandPark) ValidateParams() bool {
	return len(cmd.Args) == 2
}

func (cmd *CommandPark) Run() (string, error) {
	var output string
	sl, err := parking.Get().AddCar(*cmd.Car)
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("Allocated slot number: %d", sl.Index)
	return output, err
}

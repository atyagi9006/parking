package cmd

import (
	"fmt"
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandSlotNumberCarNumber struct {
	Command
	CarNumber string
}

func (cmd *CommandSlotNumberCarNumber) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	cmd.CarNumber = cmd.Args[0]
	return nil
}

func (cmd *CommandSlotNumberCarNumber) Clear() {
	cmd.Args = nil
	cmd.CarNumber = ""
}

func (cmd *CommandSlotNumberCarNumber) ValidateParams() bool {
	return len(cmd.Args) == 1 && cmd.Args[0] != ""
}

func (cmd *CommandSlotNumberCarNumber) Run() (string, error) {
	var output string
	slot := parking.Get().GetSlotByCarNumber(cmd.CarNumber)
	if slot == nil {
		return "Not found", nil
	}
	output = fmt.Sprint(slot.Index)
	return output, nil
}

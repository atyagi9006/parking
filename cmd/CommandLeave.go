package cmd

import (
	"fmt"
	"strconv"
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandLeave struct {
	Command
	SlotNumber uint
}

func (cmd *CommandLeave) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	value, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		return errInvalidParam
	}
	cmd.SlotNumber = uint(value)
	return nil
}

func (cmd *CommandLeave) Clear() {
	cmd.Args = nil
	cmd.SlotNumber = 0
}

func (cmd *CommandLeave) ValidateParams() bool {
	return len(cmd.Args) == 1
}

func (cmd *CommandLeave) Run() (string, error) {
	var output string
	park := parking.Get()
	slotNumber := cmd.SlotNumber
	if err := park.RemoveCarBySlot(slotNumber); err != nil {
		return output, err
	}
	output = fmt.Sprintf("Slot number %d is free", slotNumber)
	return output, nil
}

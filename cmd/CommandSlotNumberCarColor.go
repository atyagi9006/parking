package cmd

import (
	"fmt"
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandSlotNumberCarColor struct {
	Command
	CarColor string
}

func (cmd *CommandSlotNumberCarColor) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	cmd.CarColor = cmd.Args[0]
	return nil
}

func (cmd *CommandSlotNumberCarColor) Clear() {
	cmd.Args = nil
	cmd.CarColor = ""
}

func (cmd *CommandSlotNumberCarColor) ValidateParams() bool {
	return len(cmd.Args) == 1 && cmd.Args[0] != ""
}

func (cmd *CommandSlotNumberCarColor) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(cmd.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, fmt.Sprint(sl.Index))
	}
	output = strings.Join(list, ", ")
	return output, nil
}

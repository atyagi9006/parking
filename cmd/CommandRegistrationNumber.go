package cmd

import (
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandRegistrationNumber struct {
	Command
	CarColor string
}

func (cmd *CommandRegistrationNumber) ParseArgs(args string) error {
	cmd.Args = strings.Split(args, " ")
	if !cmd.ValidateParams() {
		return errInvalidParam
	}
	cmd.CarColor = cmd.Args[0]
	return nil
}

func (cmd *CommandRegistrationNumber) Clear() {
	cmd.Args = nil
	cmd.CarColor = ""
}

func (cmd *CommandRegistrationNumber) ValidateParams() bool {
	return len(cmd.Args) == 1 && cmd.Args[0] != ""
}

func (cmd *CommandRegistrationNumber) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(cmd.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, sl.GetCarNumber())
	}
	output = strings.Join(list, ", ")
	return output, nil
}

package cmd

import (
	"fmt"
	"strings"

	parking "github.com/atyagi9006/parking/lot"
)

type CommandStatus struct {
	Command
}

func (cmd *CommandStatus) ParseArgs(args string) error {
	return nil
}

func (cmd *CommandStatus) Clear() {
	cmd.Args = nil
}

func (cmd *CommandStatus) ValidateParams() bool {
	return true
}

func (cmd *CommandStatus) Run() (string, error) {
	var list = []string{fmt.Sprintf("%-12s%-20s%-10s", "Slot No.", "Registration No", "Colour")}
	filledSlots := parking.Get().GetFilledSlots()
	for _, sl := range filledSlots {
		cr := sl.Car
		list = append(list, fmt.Sprintf("%-12v%-20v%-10v", sl.Index, cr.Number, cr.Color))
	}
	output := strings.Join(list, "\n")
	return output, nil
}

package slot

import (
	"fmt"

	"github.com/atyagi9006/parking/car"
)

// Slot is a cleared area that is intended for parking car, with identity serial index number.
type Slot struct {
	Index uint
	Car   *car.Car
}

func New() *Slot {
	return &Slot{}
}

func (s *Slot) Allocate(cr car.Car) error {
	if s.Car != nil {
		return fmt.Errorf("slot: Slot already allocated")
	}
	s.Car = &cr
	return nil
}

func (s *Slot) GetCarNumber() string {
	return s.Car.Number
}

func (s *Slot) GetCarColor() string {
	return s.Car.Color
}

func (s *Slot) Free() {
	s.Car = nil
}

func (s *Slot) IsFree() bool {
	return s.Car == nil
}

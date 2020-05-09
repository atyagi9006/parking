package lot

import (
	"fmt"
	"strings"

	"github.com/atyagi9006/parking/car"
	"github.com/atyagi9006/parking/slot"
)

// Parking is main component, for data structure a parking lot
type Parking struct {
	Capacity uint
	Slots    []*slot.Slot
}

var (
	saved      *Parking
	startIndex = 1
)

func New(capacity uint) *Parking {
	parking := new(Parking)
	parking.Capacity = capacity
	parking.Slots = make([]*slot.Slot, capacity)
	idx := startIndex
	for i := range parking.Slots {
		parking.Slots[i] = &slot.Slot{Index: uint(idx)}
		idx++
	}
	parking.Save()
	return parking
}

// Get for get saved data
func Get() *Parking {
	return saved
}

func (p *Parking) Save() {
	saved = p
}

func (p *Parking) FindNearestSlot() (*slot.Slot, error) {
	for _, sl := range p.Slots {
		if sl.IsFree() {
			return sl, nil
		}
	}
	return nil, fmt.Errorf("Sorry, parking lot is full")
}

func (p *Parking) AddCar(cr car.Car) (*slot.Slot, error) {
	sl, err := p.FindNearestSlot()
	if err != nil {
		return nil, err
	}
	if err := sl.Allocate(cr); err != nil {
		return nil, err
	}
	return sl, nil
}

func (p *Parking) GetFilledSlots() (filledSlots []*slot.Slot) {
	for _, sl := range p.Slots {
		if !sl.IsFree() {
			filledSlots = append(filledSlots, sl)
		}
	}
	return
}

func (p *Parking) GetSlotsByCarColor(carColor string) (slots []*slot.Slot) {
	for _, sl := range p.Slots {
		if !sl.IsFree() {
			if strings.ToLower(sl.GetCarColor()) == strings.ToLower(carColor) {
				slots = append(slots, sl)
			}
		}
	}
	return
}

func (p *Parking) GetSlotByCarNumber(carNumber string) (slots *slot.Slot) {
	for _, sl := range p.Slots {
		if !sl.IsFree() {
			if sl.GetCarNumber() == carNumber {
				slots = sl
				return
			}
		}
	}
	return
}

func (p *Parking) RemoveCar(cr car.Car) {
	for i, sl := range p.Slots {
		if !sl.IsFree() && sl.Car.IsEqual(cr) {
			p.Slots[i].Free()
		}
	}
}

func (p *Parking) RemoveCarBySlot(slotNumber uint) error {
	for i, sl := range p.Slots {
		if sl.Index == slotNumber {
			p.Slots[i].Car = nil
			return nil
		}
	}
	return fmt.Errorf("Slot %d not found", slotNumber)
}

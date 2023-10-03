package databinding

import (
	"time"
)

type Animal struct {
	Name          string
	ArrivalDate   time.Time
	SpeciesId     int
	Speed         int
	Sex           Sex
	Weight        float64
	PreferredFood string
	Domesticated  bool
	Remarks       string
	Patience      time.Duration
}

func (a *Animal) PatienceField() *DurationField {
	return &DurationField{&a.Patience}
}

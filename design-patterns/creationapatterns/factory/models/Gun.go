package models

import "creationapatterns/factory/interfaces"

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetPower() int {
	return g.Power
}

type Ak47 struct {
	Gun
}

func NewAk47() interfaces.IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}

type Maverick struct {
	Gun
}

func NewMaverick() interfaces.IGun {
	return &Maverick{
		Gun: Gun{
			Name:  "Maverick gun",
			Power: 5,
		},
	}
}

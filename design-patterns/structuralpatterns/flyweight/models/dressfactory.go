package models

import (
	"fmt"
	"structuralpatterns/flyweight/interfaces"
)

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]interfaces.Dress),
	}
)

type DressFactory struct {
	DressMap map[string]interfaces.Dress
}

func (d *DressFactory) GetDressByType(dressType string) (interfaces.Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.DressMap[dressType] = NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

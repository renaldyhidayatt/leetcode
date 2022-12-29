package flyweight

import (
	"fmt"
	"structuralpatterns/flyweight/models"
)

func FlyWeight() {
	game := models.NewGame()
	//Add Terrorist
	game.AddTerrorist(models.TerroristDressType)
	game.AddTerrorist(models.TerroristDressType)
	game.AddTerrorist(models.TerroristDressType)
	game.AddTerrorist(models.TerroristDressType)
	//Add CounterTerrorist
	game.AddCounterTerrorist(models.CounterTerrroristDressType)
	game.AddCounterTerrorist(models.CounterTerrroristDressType)
	game.AddCounterTerrorist(models.CounterTerrroristDressType)
	dressFactoryInstance := models.GetDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}

}

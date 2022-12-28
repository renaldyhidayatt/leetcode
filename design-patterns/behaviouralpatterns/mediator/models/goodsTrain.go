package models

import (
	"behaviouralpatterns/mediator/interfaces"
	"fmt"
)

type GoodsTrain struct {
	Mediator interfaces.Mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println("Goodtrain: Landing")
	} else {
		fmt.Println("Goodtrain: Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	g.Mediator.NotifyFree()
	fmt.Println("GoodsTrain: Leaving")
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}

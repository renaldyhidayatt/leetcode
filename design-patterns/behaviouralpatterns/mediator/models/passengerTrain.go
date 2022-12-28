package models

import (
	"behaviouralpatterns/mediator/interfaces"
	"fmt"
)

type PassengerTrain struct {
	Mediator interfaces.Mediator
}

func (g *PassengerTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *PassengerTrain) Departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.NotifyFree()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}

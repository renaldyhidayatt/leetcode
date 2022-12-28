package mediator

import (
	"behaviouralpatterns/mediator/models"
	"behaviouralpatterns/mediator/services"
)

func Mediator() {
	stationManager := services.NewStationManager()

	PassengerTrain := &models.PassengerTrain{
		Mediator: stationManager,
	}
	GoodTrain := &models.GoodsTrain{
		Mediator: stationManager,
	}

	PassengerTrain.RequestArrival()
	GoodTrain.RequestArrival()
	PassengerTrain.RequestArrival()

}

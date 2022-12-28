package memento

import (
	"behaviouralpatterns/memento/models"
	"fmt"
)

func Memento() {
	caretaker := &models.CareTaker{
		MementoArray: make([]*models.Memento, 0),
	}
	originator := &models.Originator{
		State: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())

	caretaker.AddMemento(originator.CreateMemento())
	originator.SetState("C")

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.Restorememento(caretaker.GetMenento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.Restorememento(caretaker.GetMenento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())
}

package models

import (
	"behaviouralpatterns/observer/interfaces"
	"fmt"
)

type Item struct {
	ObServerList []interfaces.Observer
	Name         string
	InStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) Register(o interfaces.Observer) {
	i.ObServerList = append(i.ObServerList, o)
}

func (i *Item) Deregister(o interfaces.Observer) {
	i.ObServerList = RemoveFromslice(i.ObServerList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObServerList {
		observer.Update(i.Name)
	}
}

func RemoveFromslice(observerList []interfaces.Observer, observerToRemove interfaces.Observer) []interfaces.Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

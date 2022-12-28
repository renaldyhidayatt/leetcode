package models

import (
	"behaviouralpatterns/state/interfaces"
	"fmt"
)

type VendingMachine struct {
	HasItem       interfaces.State
	ItemRequested interfaces.State
	HasMoney      interfaces.State
	NoItem        interfaces.State

	CurrentState interfaces.State

	ItemCount int
	ItemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		VendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		VendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		VendingMachine: v,
	}
	noItemState := &NoItemState{
		VendingMachine: v,
	}

	v.SetState(hasItemState)
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}

func (v *VendingMachine) SetState(s interfaces.State) {
	v.CurrentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.ItemCount = v.ItemCount + count
}

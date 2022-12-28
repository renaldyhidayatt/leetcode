package service

import (
	"behaviouralpatterns/chainofrespon/interfaces"
	"behaviouralpatterns/chainofrespon/models"
	"fmt"
)

type Cashier struct {
	next interfaces.Department
}

func (c *Cashier) Execute(p *models.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next interfaces.Department) {
	c.next = next
}

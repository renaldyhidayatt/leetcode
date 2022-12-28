package service

import (
	"behaviouralpatterns/chainofrespon/interfaces"
	"behaviouralpatterns/chainofrespon/models"
	"fmt"
)

type Reception struct {
	next interfaces.Department
}

func (r *Reception) Execute(p *models.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next interfaces.Department) {
	r.next = next
}

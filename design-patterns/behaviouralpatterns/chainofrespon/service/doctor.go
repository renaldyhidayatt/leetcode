package service

import (
	"behaviouralpatterns/chainofrespon/interfaces"
	"behaviouralpatterns/chainofrespon/models"
	"fmt"
)

type Doctor struct {
	next interfaces.Department
}

func (d *Doctor) Execute(p *models.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
	}

	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next interfaces.Department) {
	d.next = next
}

package service

import (
	"behaviouralpatterns/chainofrespon/interfaces"
	"behaviouralpatterns/chainofrespon/models"
	"fmt"
)

type Medical struct {
	next interfaces.Department
}

func (m *Medical) Execute(p *models.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}

	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next interfaces.Department) {
	m.next = next
}

package chainofrespon

import (
	"behaviouralpatterns/chainofrespon/models"
	"behaviouralpatterns/chainofrespon/service"
)

func ChainOfRespon() {
	cashier := &service.Cashier{}
	//Set next for medical department
	medical := &service.Medical{}
	medical.SetNext(cashier)
	//Set next for doctor department
	doctor := &service.Doctor{}
	doctor.SetNext(medical)
	//Set next for reception department
	reception := &service.Reception{}
	reception.SetNext(doctor)
	patient := &models.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}

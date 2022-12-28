package interfaces

import "behaviouralpatterns/chainofrespon/models"

type Department interface {
	Execute(*models.Patient)
	SetNext(Department)
}

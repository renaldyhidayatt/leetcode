package interfaces

import "behaviouralpatterns/iterator/models"

type Iterator interface {
	HasNext() bool
	GetNext() *models.User
}

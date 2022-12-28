package services

import (
	"behaviouralpatterns/iterator/interfaces"
	"behaviouralpatterns/iterator/models"
)

type UserCollection struct {
	Users []*models.User
}

func (u *UserCollection) CreateIterator() interfaces.Iterator {
	userIter := &models.UserIterator{
		Users: u.Users,
	}
	return userIter
}

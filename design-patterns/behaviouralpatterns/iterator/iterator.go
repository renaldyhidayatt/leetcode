package iterator

import (
	"behaviouralpatterns/iterator/models"
	"behaviouralpatterns/iterator/services"
	"fmt"
)

func Iterator() {
	user1 := &models.User{
		Name: "a",
		Age:  30,
	}
	user2 := &models.User{
		Name: "b",
		Age:  20,
	}
	userCollection := &services.UserCollection{
		Users: []*models.User{user1, user2},
	}
	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

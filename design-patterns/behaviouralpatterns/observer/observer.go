package observer

import "behaviouralpatterns/observer/models"

func ObServer() {
	shirtItem := models.NewItem("Nike Shirt")
	observerFirst := &models.Customer{Id: "abc@gmail.com"}
	observerSecond := &models.Customer{Id: "xyz@gmail.com"}
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}

package getgun

import (
	"creationapatterns/factory/interfaces"
	"creationapatterns/factory/models"
	"fmt"
)

func GetGun(gunType string) (interfaces.IGun, error) {
	if gunType == "ak47" {
		return models.NewAk47(), nil
	}
	if gunType == "maverick" {
		return models.NewMaverick(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

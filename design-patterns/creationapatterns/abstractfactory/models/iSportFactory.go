package models

import (
	"creationapatterns/abstractfactory/interfaces"
	"creationapatterns/abstractfactory/schemas"
	"fmt"
)

type ISportsFactory interface {
	MakeShoe() interfaces.IShoe
	MakeShort() interfaces.IShort
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &schemas.Adidas{}, nil
	}
	if brand == "nike" {
		return &schemas.Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

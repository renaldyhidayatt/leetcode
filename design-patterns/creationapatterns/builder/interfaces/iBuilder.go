package interfaces

import "creationapatterns/builder/schemas"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() schemas.House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &schemas.NormalBuilder{}
	}
	if builderType == "igloo" {
		return &schemas.IglooBuilder{}
	}
	return nil
}

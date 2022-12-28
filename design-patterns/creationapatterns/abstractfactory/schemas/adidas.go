package schemas

import "creationapatterns/abstractfactory/interfaces"

type Adidas struct {
}

func (a *Adidas) MakeShoe() interfaces.IShoe {
	return &adidasShoe{
		interfaces.Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShort() interfaces.IShort {
	return &adidasShort{
		interfaces.Short{
			Logo: "adidas",
			Size: 14,
		},
	}
}

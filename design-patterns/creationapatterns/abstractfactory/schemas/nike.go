package schemas

import "creationapatterns/abstractfactory/interfaces"

type Nike struct {
}

func (a *Nike) MakeShoe() interfaces.IShoe {
	return &nikeShoe{
		interfaces.Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Nike) MakeShort() interfaces.IShort {
	return &nikeShort{
		interfaces.Short{
			Logo: "adidas",
			Size: 14,
		},
	}
}

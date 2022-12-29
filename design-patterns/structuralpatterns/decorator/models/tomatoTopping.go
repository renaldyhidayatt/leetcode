package models

import "structuralpatterns/decorator/interfaces"

type TomatoTopping struct {
	Pizza interfaces.Pizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}

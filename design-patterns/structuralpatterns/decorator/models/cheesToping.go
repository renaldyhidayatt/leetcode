package models

import "structuralpatterns/decorator/interfaces"

type CheeseTopping struct {
	Pizza interfaces.Pizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}

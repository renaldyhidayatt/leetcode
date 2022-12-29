package decorator

import (
	"fmt"
	"structuralpatterns/decorator/models"
)

func Decorator() {
	veggiePizza := &models.VeggeMania{}

	//Add cheese topping
	veggiePizzaWithCheese := &models.CheeseTopping{
		Pizza: veggiePizza,
	}

	//Add tomato topping
	veggiePizzaWithCheeseAndTomato := &models.TomatoTopping{
		Pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.GetPrice())

	peppyPaneerPizza := &models.PeppyPaneer{}

	//Add cheese topping
	peppyPaneerPizzaWithCheese := &models.CheeseTopping{
		Pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.GetPrice())
}

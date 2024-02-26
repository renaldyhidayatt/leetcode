package lasagnamaster

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer < 1 {
		return len(layers) * 2
	}

	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendList []string, myList []string) {
	(myList)[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	newQuantities := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		newQuantities[i] = quantities[i] * float64(numberOfPortions) / 2
	}

	return newQuantities
}

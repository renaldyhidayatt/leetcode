package animalmagic

import (
	"math/rand"
	"time"
)

func RollADie() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(20) + 1
}

func GenerateWandEnergy() float64 {
	rand.Seed(time.Now().UnixNano())

	return rand.Float64() * 12.0
}

func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}

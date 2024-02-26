package carsassemble

func CalculateWorkingCarsPerHours(productionRate int, successRate float64) float64 {
	successfulCars := float64(productionRate) * (successRate / 100)

	return successfulCars
}

func CalculateWorkingCarsPerMinutes(productionRate int, successRate float64) int {
	successFulCarsPerHour := CalculateWorkingCarsPerHours(productionRate, successRate)

	successFulCarsPerMinutes := successFulCarsPerHour / 60

	return int(successFulCarsPerMinutes)
}

func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10

	remainingCars := carsCount % 10
	totalCost := uint(groupsOfTen*95000 + remainingCars*10000)

	return totalCost

}

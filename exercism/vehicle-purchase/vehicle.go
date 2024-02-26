package vehiclepurchase

func NeedLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}

	return option2 + " is clearly the better choice."
}

func CalculateResllPrice(originalPrice, age float64) float64 {
	var resellPrice float64

	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age >= 10 {
		resellPrice = originalPrice * 0.5
	} else {
		resellPrice = originalPrice * 0.7
	}

	return resellPrice
}

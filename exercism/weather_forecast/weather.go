package weatherforecast

var CurrentCondition string

var CurrentLocation string

func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	return CurrentCondition + " - current weather condition: " + CurrentCondition
}

// Package weather provides functionality for generating a weather forcast.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation describes the current location for the forecast.
var CurrentLocation string

// Forecast generates a weather forecast based on the current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

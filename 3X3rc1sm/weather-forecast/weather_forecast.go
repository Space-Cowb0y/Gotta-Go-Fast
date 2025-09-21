// Package weather provides weather information.
package weather

// CurrentCondition holds the current weather condition as a string (e.g., "Sunny", "Rainy").
var CurrentCondition string

// CurrentLocation holds the name of the current location as a string (e.g., "San Francisco", "New York").
var CurrentLocation string

// Forecast returns a string with the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

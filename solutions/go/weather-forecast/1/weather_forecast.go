// packages are the name
package weather

// CurrentCondition is the weather condition in the current location
var CurrentCondition string
var CurrentLocation string

// Forecast returns a string with the current location and weather condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

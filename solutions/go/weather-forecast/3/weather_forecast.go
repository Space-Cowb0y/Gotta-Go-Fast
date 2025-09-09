// The package comment should introduce the package and provide information relevant to the package as a whole.
package weather

// This should tell any user of the package what information the variables store, and what they can do with it.
// CurrentCondition holds the current weather condition as a string (e.g., "Sunny", "Rainy").
// CurrentLocation holds the name of the current location as a string (e.g., "San Francisco", "New York").
var CurrentCondition string
var CurrentLocation string

// Forecast returns a string with the current location and weather condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

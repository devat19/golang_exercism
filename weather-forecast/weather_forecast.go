/*
Package weather returns the city's current conditions in string format.
*/
package weather

// CurrentCondition - String data type.
var CurrentCondition string

// CurrentLocation - String data type.
var CurrentLocation string

// Forecast function returns city's forecast in string format.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

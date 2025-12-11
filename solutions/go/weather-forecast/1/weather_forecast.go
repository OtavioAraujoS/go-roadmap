// Package weather provides tools to store and share weather forecasts
// for different locations in Goblinocus.
package weather

// CurrentCondition stores the latest reported weather condition
// for the currently selected location.
var CurrentCondition string

// CurrentLocation stores the name of the location for which the current
// weather condition has been reported.
var CurrentLocation string

// Forecast provides the weather forecast for a given city by updating
// the current location and weather condition and returning a formatted summary.
func Forecast(city, condition string) string {
    CurrentLocation = city
    CurrentCondition = condition
    return city + " - current weather condition: " + condition
}

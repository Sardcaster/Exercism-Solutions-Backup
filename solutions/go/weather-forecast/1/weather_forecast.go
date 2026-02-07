// Package weather is a func.
package weather 

var (
    // CurrentCondition something.
	CurrentCondition string
    // CurrentLocation something.
	CurrentLocation  string
)

// Forecast() something.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

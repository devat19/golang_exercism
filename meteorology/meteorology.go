package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (sc TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[sc]
}

func (d Temperature) String() string {
	return fmt.Sprintf("%v %v", d.degree, d.unit)
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (sc SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sc]
}

func (d Speed) String() string {
	return fmt.Sprintf("%v %v", d.magnitude, d.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (d MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %v%% Humidity", d.location, d.temperature.String(), d.windDirection, d.windSpeed.String(), d.humidity)
}

// Add a String method to MeteorologyData

package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
	// panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var lexicographicalVehicleChoice string = ""

	if option1 <= option2 {
		lexicographicalVehicleChoice = option1
	} else {
		lexicographicalVehicleChoice = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", lexicographicalVehicleChoice)
	// panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	}
	if age < 10 {
		return originalPrice * 0.7
	}
	return originalPrice * 0.5
	// panic("CalculateResellPrice not implemented")
}

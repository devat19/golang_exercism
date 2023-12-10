package elon

import "fmt"

// TODO: define the 'Drive()' method

func (c *Car) Drive() {

	carBatteryAvailable := c.battery
	chargePercentageNeeded := c.batteryDrain

	var carBatteryPostDrive int = int(carBatteryAvailable) - int(chargePercentageNeeded)

	if carBatteryPostDrive >= 0 {
		c.battery = c.battery - chargePercentageNeeded
		c.distance = c.distance + c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
// Understand the difference between -> calling it by * (pointer). Can these be used concurrently
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	// return fmt.Sprintf("Battery at %v%", c.battery)
	return "Battery at " + fmt.Sprint(c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

func (c *Car) CanFinish(trackDistance int) bool {
	// return fmt.Sprintf("Battery at %v%", c.battery)
	var maxDistance float32 = float32(c.battery/c.batteryDrain) * float32(c.speed)
	return maxDistance-float32(trackDistance) >= 0
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

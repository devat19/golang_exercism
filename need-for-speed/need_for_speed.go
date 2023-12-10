package speed

import "fmt"

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	carBatteryAvailable := car.battery
	chargePercentageNeeded := car.batteryDrain

	var carBatteryPostDrive int = int(carBatteryAvailable) - int(chargePercentageNeeded)

	if carBatteryPostDrive < 0 {
		return car
	} else {
		car.battery = car.battery - chargePercentageNeeded
		car.distance = car.distance + car.speed
		return car
	}

	// panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	/*
		td := float64(track.distance)
	    carTimes := int(math.Ceil(td/float64(car.speed)))
	    carDrain := car.batteryDrain * carTimes
	    return carDrain <= car.battery
	*/
	var maxDistance float32 = float32(car.battery/car.batteryDrain) * float32(car.speed)

	fmt.Println("data", maxDistance, track.distance, car, maxDistance-float32(track.distance))

	// return (maxDistance <= float32(track.distance))

	// if maxDistance-float32(track.distance) < 0 {
	// 	return false
	// }
	// return true
	return maxDistance-float32(track.distance) >= 0

	// if maxDistance <= float32(track.distance) {
	// 	return true
	// }
	// return false
}

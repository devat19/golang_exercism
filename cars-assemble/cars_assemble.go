package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// Typecasting!
	var cst float64 = (float64(productionRate) * successRate / 100)
	return cst
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return (int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	// Different ways of defining a variable in GO. - explicit and implicit

	// var carsCountInABundle int = 10
	// var productionCostPerBundle int = 95000 // Dollars
	// var productionCostPerCar int = 10000    // Dollars

	// carsCountInABundle, productionCostPerBundle, productionCostPerCar := 10, 95000, 10000

	// const carsCountInABundle int = 10
	// const productionCostPerBundle int = 95000 // Dollars
	// const productionCostPerCar int = 10000    // Dollars

	carsCountInABundle := 10
	productionCostPerBundle := 95000 // Dollars
	productionCostPerCar := 10000    // Dollars

	var bundleCount int = ((carsCount) / carsCountInABundle)
	var totalBundledCost int = bundleCount * productionCostPerBundle
	var remainingCarsCost int = ((carsCount) % carsCountInABundle) * productionCostPerCar

	return uint(totalBundledCost + remainingCarsCost)
}

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const costPerIndividually = 10000
	const costPerGroupOfTen = 95000

	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10

	totalCost := (groupsOfTen * costPerGroupOfTen) + (remainingCars * costPerIndividually)
	return uint(totalCost)	
}

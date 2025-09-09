package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var rate float64 = float64(productionRate) * (successRate / 100)
	return rate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var cost uint
	if carsCount < 1 {
		cost = 0
	} else if carsCount <= 9 {
		cost = uint(carsCount) * 10000
	} else if carsCount >= 10 {
		cost = uint(carsCount) * 9500
	} else {
		cost = uint(carsCount)*9000 + 1000
	}
	return cost
}

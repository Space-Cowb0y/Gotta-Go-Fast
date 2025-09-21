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
	//10 cars cpst 95,000
	//9 cars cost 90,000
	//1-8 cars cost 10,000 each
	if carsCount >= 10 {
		cost += uint((carsCount / 10) * 95000)
		carsCount = carsCount % 10
	}
	if carsCount <= 9 {
		cost += uint(carsCount * 10000)
		carsCount -= 9
	}
	return cost
}

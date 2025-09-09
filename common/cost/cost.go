package cost

func GetBaseRate(volume int) int {
	const lowRate = 800
	const mediumRate = 1200
	const highRate = 1500

	if volume <= 6000 {
		return lowRate
	} else if volume > 6000 && volume < 100000 {
		return mediumRate
	} 
	return highRate
}



func GetShippingRate(baseRate int, island int) int {
	const northIsland = 1
	const northIslandRate = 10
	const southIsland = 2
	const southIslandRate = 15
	const stewartIsland = 3
	const stewartIslandRate = 20
	
	if island == northIsland {
		return baseRate * northIslandRate
	} else if island == southIsland {
		return baseRate * southIslandRate
	}
	return baseRate * stewartIslandRate
}
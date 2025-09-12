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



func GetShippingRate(baseRate int, island string) int {

	const northIslandRate = 10
	
	const southIslandRate = 15
	
	const stewartIslandRate = 20
	
  switch island {
    case "northIsland":
        return baseRate * northIslandRate
    case "southIsland":
        return baseRate * southIslandRate
	case "stewartIsland":
		baseRate * stewartIslandRate
	default:
        return 
    }
}
package cost

import "strings"

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
	
	isNorthIsland := strings.Contains(strings.ToLower(island), "north")
	isSouthIsland := strings.Contains(strings.ToLower(island), "south")
	isStewartIsland := strings.Contains(strings.ToLower(island), "stewart")

  switch {
    case isNorthIsland:
        return baseRate * northIslandRate
    case isSouthIsland:
        return baseRate * southIslandRate
	case isStewartIsland:
		return baseRate * stewartIslandRate
	default:
        return baseRate
    }
}
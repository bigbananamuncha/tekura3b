package cost

import (
	"strings"
)

//handling invalid island inputs
func IsValidIsland(island string) bool {
    normalized := strings.ToLower(island)
    return strings.Contains(normalized, "north") ||
           strings.Contains(normalized, "south") ||
           strings.Contains(normalized, "stewart")
}

//Find the base rate of the shipping cost due to the volume of the package 
func GetBaseRate(volume float64) float64 {
	const lowRate = 800
	const mediumRate = 1200
	const highRate = 1500

	if volume <= 6000 {
		return lowRate
	} else if volume > 6000 && volume <= 100000 {
		return mediumRate
	} 
	return highRate
}

//find the final shipping rate using the surcharge from each island
func GetShippingRate(baseRate float64, island string) float64 {

	var finalRate float64 = baseRate
	const northIslandRate = 1.10
	const southIslandRate = 1.15
	const stewartIslandRate = 1.20
	
	isNorthIsland := strings.Contains(strings.ToLower(island), "north")
	isSouthIsland := strings.Contains(strings.ToLower(island), "south")
	isStewartIsland := strings.Contains(strings.ToLower(island), "stewart")

  switch {
    case isNorthIsland:
    	finalRate = baseRate * northIslandRate
    case isSouthIsland:
        finalRate = baseRate * southIslandRate
	case isStewartIsland:
		finalRate = baseRate * stewartIslandRate
	default:
        finalRate = baseRate
    }
	return finalRate
}
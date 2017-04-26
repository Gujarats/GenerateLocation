package location

import "fmt"

// This const is used for generating the latitdue and longitude
// only used in newPoint() function
const (
	// this number is aprroximate from 1 seconds to meters
	meters = 24.384
)

type Location struct {
	Lat float64
	Lon float64
}

// Generate new location from the given point to east,
// and repeat it until specific of length in km.
// so if the limitLeght is 40 km then the generate location would be 40 km to east and 40 km to south.
// note that the given latitude and longitude must be in the left top of the square.
// separate and add new location with given distance in km addition
// NOTE : distance and limitLength must be in km
func GenerateLocation(lat, lon float64, distance int, limitLength int) []Location {
	// create array location for storing the location
	var locations []Location

	// Generate location to East
	for counterDistanceEast := distance; counterDistanceEast <= limitLength; counterDistanceEast += distance {
		newLatEast, newLonEast := newPoint(lat, lon, counterDistanceEast, "east")
		locations = append(locations, Location{Lat: newLatEast, Lon: newLonEast})
	}

	// looping locationEast to Generate location South
	for _, locationEast := range locations {
		for counterDistanceSouth := distance; counterDistanceSouth < limitLength; counterDistanceSouth += distance {
			newLatSouth, newLonSouth := newPoint(locationEast.Lat, locationEast.Lon, counterDistanceSouth, "south")
			locations = append(locations, Location{Lat: newLatSouth, Lon: newLonSouth})
		}

	}

	return locations

}

// Get the center Location
func GetCenterLocation(lat, lon float64, distance int, limitLength int) Location {
	var locations []Location
	baseCenter := (limitLength / distance) / 2
	fmt.Println("baseCenter = ", baseCenter)

	// Generate location to East
	for counterDistanceEast := distance; counterDistanceEast <= limitLength; counterDistanceEast += distance {
		newLatEast, newLonEast := newPoint(lat, lon, counterDistanceEast, "east")
		locations = append(locations, Location{Lat: newLatEast, Lon: newLonEast})
	}
	fmt.Println("Location length = ", len(locations))

	// looping locationEast to Generate location South
	for indexEast, locationEast := range locations {
		indexSouth := 0
		for counterDistanceSouth := distance; counterDistanceSouth < limitLength; counterDistanceSouth += distance {
			indexSouth++
			//fmt.Println("indexSouth = ", indexSouth)
			//fmt.Println("indexEast = ", indexEast)
			if indexEast+1 == baseCenter && indexSouth+1 == baseCenter {
				newLatSouth, newLonSouth := newPoint(locationEast.Lat, locationEast.Lon, counterDistanceSouth, "south")
				return Location{
					Lat: newLatSouth,
					Lon: newLonSouth,
				}
			}
		}

	}

	return Location{}
}

// distance must be in km
// direction could be west,east,north,south
func newPoint(lat, lon float64, distance int, direction string) (float64, float64) {
	// conver distance to meters
	// we need to convert it to meters because this will be divided by 1 seconds or 24 in meters
	distanceMeters := float64(distance * 1000.0)

	// get seconds
	seconds := distanceMeters / meters

	//convert seconds to decimal
	additionalDecimal := secondsToDecimal(seconds)

	switch direction {
	case "west":
		//gives negative
		lon = lon - additionalDecimal
	case "east":
		lon = lon + additionalDecimal
	case "north":
		lat = lat + additionalDecimal
	case "south":
		// gives negative
		lat = lat - additionalDecimal
	default:
		fmt.Errorf("Given direction is not available")
		return lat, lon
	}

	return lat, lon

}

// convert seconds to decimal
func secondsToDecimal(seconds float64) float64 {
	return seconds / (60.0 * 60.0)
}

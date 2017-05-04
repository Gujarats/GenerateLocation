package location

import "fmt"

// This const is used for generating the latitdue and longitude
// only used in newPoint() function
const (
	// this number is aprroximate from 1 seconds to meters
	meters = 24.384
)

// This one is used for getting the center of the marked location by splitting them using quadran.
type centerLocation struct {
	QuadranLevelDetail string
	CenterLocation     Location
}

// for return value GetCenterQuadranLocations() .
// the return value is ordered from index 0 to the latest.
// index 0 = center location level 0
// index 1 = center location quadran1 level 1
// index 2 = center location quadran2 level 1
// index 3 = center location quadran3 level 1
// index 4 = center location quadran4 level 1
// index 5 = center location quadran1 level 2. And so on.
var centerLocations []centerLocation

// Generate new location from the given point to east,
// and repeat it until specific of length in km.
// so if the limitLeght is 40 km then the generate location would be 40 km to east and 40 km to south.
// note that the given latitude and longitude must be in the left top of the square.
// separate and add new location with given distance in km addition
// NOTE : distance and limitLength must be in km
func (l *Location) GenerateLocation(distance, limitLength float64) []Location {
	// create array location for storing the location
	var locations []Location

	// Generate location to East
	for counterDistanceEast := distance; counterDistanceEast <= limitLength; counterDistanceEast += distance {
		newLatEast, newLonEast := newPoint(l.Lat, l.Lon, counterDistanceEast, "east")
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

// Getting the center marked location by split the map using quadran
// For example if we have marked the location in some location it will have square shape.
// And from generated locations we created two dimensional array that will store all the locations
// We can get the center of the location by diving the square shape to half.
// Also getting another center location by its level, like center location in quadran level 1.
func (l *Location) GetCenterQuadranLocations() {

}

// Get the center Location.
// This function is working like GenererateLocation, but only get the center Location.
// Imagine that you have a square and inside that square thre are many Location, this will return the center of location inside that square.
func GetCenterLocation(lat, lon, distance, limitLength float64) Location {
	var locations []Location
	baseCenter := int((limitLength / distance) / 2)

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
func newPoint(lat, lon, distance float64, direction string) (float64, float64) {
	// conver distance to meters
	// we need to convert it to meters because this will be divided by 1 seconds or 24 in meters
	distanceMeters := distance * 1000.0

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

package location

import (
	"errors"
	"fmt"
	"math"
)

// This const is used for generating the latitdue and longitude
// only used in newPoint() function
const (
	// this number is aprroximate from 1 seconds to meters
	meters = 24.384
)

// This one is used for getting the center of the marked location by splitting them using quadran.
type CenterLocation struct {
	Level          int
	Quadran        int
	MarkedLocation Location
}

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

	// create multidimentional array to mapping the location
	//mappingLocations = make([][]Location, len(locations))
	//for index, _ := range mappingLocations {
	//	mappingLocations[index] = make([]Location, len(locations))
	//}

	//initilize value mappingLocations
	//for i := 0; i < len(locations); i++ {
	//	mappingLocations[0][i] = locations[i]
	//}

	// looping locationEast to Generate location South
	for _, locationEast := range locations {
		for counterDistanceSouth := distance; counterDistanceSouth < limitLength; counterDistanceSouth += distance {
			// generate Location to the south
			newLatSouth, newLonSouth := newPoint(locationEast.Lat, locationEast.Lon, counterDistanceSouth, "south")
			location := Location{Lat: newLatSouth, Lon: newLonSouth}

			// store our generated location
			//mappingLocations[rowIndex][columnIndex] = location
			locations = append(locations, location)
		}

	}

	return locations

}

// Give two dimension array of locations
func (l *Location) GetMultiLocations(distance, limitLength float64) [][]Location {
	// mapping the locations using two dimensional
	var mappingLocations [][]Location

	// create array location for storing the location
	var locations []Location

	// Generate location to East
	for counterDistanceEast := distance; counterDistanceEast <= limitLength; counterDistanceEast += distance {
		newLatEast, newLonEast := newPoint(l.Lat, l.Lon, counterDistanceEast, "east")
		locations = append(locations, Location{Lat: newLatEast, Lon: newLonEast})
	}

	// create multidimentional array to mapping the location
	mappingLocations = make([][]Location, len(locations))
	for index, _ := range mappingLocations {
		mappingLocations[index] = make([]Location, len(locations))
	}

	//initilize value mappingLocations
	for i := 0; i < len(locations); i++ {
		mappingLocations[0][i] = locations[i]
	}

	// looping locationEast to Generate location South
	for rowIndex, locationEast := range locations {
		columnIndex := 0
		for counterDistanceSouth := distance; counterDistanceSouth < limitLength; counterDistanceSouth += distance {
			// generate Location to the south
			newLatSouth, newLonSouth := newPoint(locationEast.Lat, locationEast.Lon, counterDistanceSouth, "south")
			location := Location{Lat: newLatSouth, Lon: newLonSouth}

			// store our generated location
			mappingLocations[rowIndex][columnIndex] = location
			locations = append(locations, location)
			columnIndex++
		}

	}

	return mappingLocations

}

// Getting the center marked location by split the map using quadran
// For example if we have marked the location in some location it will have square shape.
// And from generated locations we created two dimensional array that will store all the locations
// We can get the center of the location by divided the square shape to half.
// Also getting another center location by its level, like center location in quadran level 1.
// the return value is map first is the level and second is the array of center locations.
func (l *Location) GetCenterQuadranLocations(distance, limitLength float64, deepLevel int) (map[int][4]Location, error) {

	// getting two dimension array of location
	multiLocations := l.GetMultiLocations(distance, limitLength)

	// initialize return value using map.
	// int difine its level
	// []CenterLocation is going to be storing the center locations
	var centerLocations = make(map[int][4]Location)

	// check if the level if possible to getting the data
	if (float64(len(multiLocations)) / math.Pow(2.0, float64(deepLevel+1))) < 1.0 {
		err := errors.New("Level is not possible to get the data from given locations, please check distance and limitLength")
		return centerLocations, err
	}

	level := 0
	center := len(multiLocations) / 2 // center level 0
	var markedCenters = [][2]int{{center, center}}
	for level <= deepLevel {
		if level > 0 {

			var newMarkedCenters = [][2]int{}
			var locations = [4]Location{}
			//getting all the center quadran from the

			fmt.Printf("markedCenters = %+v\n", markedCenters)
			for _, markedCenter := range markedCenters {
				// get positions for getting all 4 center locations from its quadran
				position := float64(len(multiLocations)) / math.Pow(2.0, float64(level+1))

				// getting the locations from each quadran. in this case we have 4 quadran
				newMarkedCenters, locations = getCenterLocations(multiLocations, markedCenter, int(position))
				centerLocations[level] = locations
			}

			// set nil to markedCenters and assign to new markedCenters
			markedCenters = newMarkedCenters

		} else {
			// gave to map an array that contains only one CenterLocations which is center0
			centerLocations[level] = [4]Location{multiLocations[center][center]}
		}
		level += 1
	}

	return centerLocations, nil
}

// get all center locations from all quadran, this would be 4 locations.
func getCenterLocations(multiLocations [][]Location, centerIndex [2]int, position int) ([][2]int, [4]Location) {
	var locations [4]Location
	var markedCenters [][2]int

	// coordinate to  mapping 2d array of multiLocations
	x := centerIndex[0]
	y := centerIndex[1]

	// get quadran 1 center location
	// go left and go up
	fmt.Printf("index original = %v,%v\n", x, y)
	fmt.Printf("index = %v,%v\n", x-position, y+position)
	fmt.Printf("MultiLocations length = %+v\n", len(multiLocations))
	fmt.Printf("result = %+v\n", multiLocations[x-position][y+position])
	fmt.Println("=================================================================")
	//markedCenters[0] = [2]int{x - position, y + position}
	markedCenters = append(markedCenters, [2]int{x - position, y + position})
	locations[0] = multiLocations[x-position][y+position]

	// get quadran 2
	// go righ go up
	centerLocQ2 := multiLocations[x-position][y-position]
	//markedCenters[1] = [2]int{x - position, y - position}
	markedCenters = append(markedCenters, [2]int{x - position, y - position})
	locations[1] = centerLocQ2

	// get quadran 3
	// go down go left
	centerLocQ3 := multiLocations[x+position][y-position]
	//markedCenters[2] = [2]int{x + position, y - position}
	markedCenters = append(markedCenters, [2]int{x + position, y - position})
	locations[2] = centerLocQ3

	// getquadra 4
	// go down go right
	centerLocQ4 := multiLocations[x+position][y+position]
	//markedCenters[3] = [2]int{x + position, y + position}
	markedCenters = append(markedCenters, [2]int{x + position, y + position})
	locations[3] = centerLocQ4

	return markedCenters, locations
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

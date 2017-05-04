//	Generate Latitude and Longitude from given location.
// To use this package first create Location struct using New() function and give the appropiate argument.
package location

import "math/rand"

// This struct is used for get base location like latitude and longitude.
type Location struct {
	// Lat,Lon for longitude and latitude.
	Lat float64
	Lon float64
}

// To store Format the latitude and longitude to Degree
type LocationDegree struct {
	Degree  int
	Minutes int
	Seconds int
}

func New(lat, lon float64) *Location {
	return &Location{
		Lat: lat,
		Lon: lon,
	}
}

// Convert the latitude and longitude into (Degre Seconds mninute) format.
// The return value of latitude and longitude in order
func (l *Location) ToDegree() (LocationDegree, LocationDegree) {
	lat := toDegree(l.Lat)
	lon := toDegree(l.Lon)

	return lat, lon

}

// convert given location Value into Decimal.
func toDecimal(locationDegree LocationDegree) float64 {
	totalSeconds := locationDegree.Degree * 60 * 60
	totalSeconds += locationDegree.Minutes * 60
	totalSeconds += locationDegree.Seconds
	return float64(totalSeconds) / (60.0 * 60.0)
}

// add seconds to lat or lon so we can get new location.
// we created array location so we can randomly select lat and lon to add new seconds degree.
// getting seconds value using pseu-do random from given seconds.
// returning latitude and longitude in order.
func (l *Location) RandomLatLongSeconds(seconds int) (float64, float64) {
	location := []float64{l.Lat, l.Lon}
	randomIndex := rand.Intn(len(location))

	randomSeconds := getRandomNumber(seconds)

	// convert the lat or lon depends on random selection to degree
	degreeLocation := toDegree(location[randomIndex])
	degreeLocation.Seconds += randomSeconds

	// convert the degree to decimal degree and replace it to the current index
	location[randomIndex] = toDecimal(degreeLocation)

	// return lat and lon in orner
	return location[0], location[1]
}

// add minute to lat or lon so we can get new location.
// we created array location so we can randomly select lat and lon to add new minute degree.
// getting minute value using pseu-do random from given minute.
// returning latitude and longitude in order.
func (l *Location) RandomLatLongMinute(minute int) (float64, float64) {

	location := []float64{l.Lat, l.Lon}
	randomIndex := rand.Intn(len(location))

	randomMinute := getRandomNumber(minute)

	// convert the lat or lon depends on random selection to degree
	degreeLocation := toDegree(location[randomIndex])
	degreeLocation.Minutes += randomMinute

	// convert the degree to decimal degree and replace it to the current index
	location[randomIndex] = toDecimal(degreeLocation)

	// return lat and lon in orner
	return location[0], location[1]
}

// Get Random number from the maximal number given
// minimal number is 1 to max number
func getRandomNumber(max int) int {
	return rand.Intn(max-1) + 1
}

// give location as an argument to function as the latitude or longitude.
// this function will convert the given Latitude or longitude depends on what you give on the location to Degree.
func toDegree(location float64) LocationDegree {
	// get the location in seconds, rounded up if needed
	seconds := int(location*60.0*60.0 + 0.5)
	minutes := seconds / 60
	degrees := seconds / (60 * 60)

	return LocationDegree{
		Degree:  degrees,
		Minutes: minutes % 60,
		Seconds: seconds % 60,
	}

}

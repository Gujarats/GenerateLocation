// How to use : first call SetupLocation(InputLat,InputLon float64) to set base latitude and longitude.
// To genreate the location you can call RandomLatLong() to generate using seconds degree.
// And RandomLatLongMinute to genereate new location using minute degree.
package location

import "math/rand"

var lat, lon float64

type LocationDegree struct {
	Degree  int
	Minutes int
	Seconds int
}

// setup base location for generating the new randon lat, lon value.
func SetupLocation(InputLat, InputLon float64) {
	lat = InputLat
	lon = InputLon
}

// Convert the latitude or longitudet into Degre Seconds mninute
func ToDegree(location float64) LocationDegree {
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

// convert given location Value into Decimal.
func ToDecimal(locationDegree LocationDegree) float64 {
	totalSeconds := locationDegree.Degree * 60 * 60
	totalSeconds += locationDegree.Minutes * 60
	totalSeconds += locationDegree.Seconds
	return float64(totalSeconds) / (60.0 * 60.0)
}

// add seconds to lat or lon so we can get new location.
// we created array location so we can randomly select lat and lon to add new seconds degree.
// getting seconds value using pseu-do random from given seconds.
// returning latitude and longitude in order.
func RandomLatLong(seconds int) (float64, float64) {
	location := []float64{lat, lon}
	randomIndex := rand.Intn(len(location))

	randomSeconds := getRandomNumber(seconds)

	// convert the lat or lon depends on random selection to degree
	degreeLocation := ToDegree(location[randomIndex])
	degreeLocation.Seconds += randomSeconds

	// convert the degree to decimal degree and replace it to the current index
	location[randomIndex] = ToDecimal(degreeLocation)

	// return lat and lon in orner
	return location[0], location[1]
}

// add minute to lat or lon so we can get new location.
// we created array location so we can randomly select lat and lon to add new seconds degree.
// getting minute value using pseu-do random from given minute.
// returning latitude and longitude in order.
func RandomLatLongMinute(minute int) (float64, float64) {

	location := []float64{lat, lon}
	randomIndex := rand.Intn(len(location))

	randomMinute := getRandomNumber(minute)

	// convert the lat or lon depends on random selection to degree
	degreeLocation := ToDegree(location[randomIndex])
	degreeLocation.Minutes += randomMinute

	// convert the degree to decimal degree and replace it to the current index
	location[randomIndex] = ToDecimal(degreeLocation)

	// return lat and lon in orner
	return location[0], location[1]
}

// Get Random number from the maximal number given
// minimal number is 1 to max number
func getRandomNumber(max int) int {
	return rand.Intn(max-1) + 1
}

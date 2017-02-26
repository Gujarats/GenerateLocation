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

func ToDegree(location float64) LocationDegree {
	// get degree
	degree := int(location)
	location = location - float64(degree)

	// get minute
	multiply := location * 60.0
	minutes := int(multiply)
	location = multiply - float64(minutes)

	// get seconds
	multiply = location * 60
	seconds := int(multiply)

	// check if the decimal is 9 or 8 we are gointo round up
	location = multiply - float64(seconds)
	if location > 0.5 {
		seconds += 1
	}

	return LocationDegree{
		Degree:  degree,
		Minutes: minutes,
		Seconds: seconds,
	}

}

func ToDecimal(locationDegree LocationDegree) float64 {
	var result float64
	result = float64(locationDegree.Seconds) / 60.0

	result += float64(locationDegree.Minutes)

	result = result / 60.0

	result += float64(locationDegree.Degree)

	return result
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

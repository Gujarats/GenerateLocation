# Generate Random Location [![Build Status](https://secure.travis-ci.org/Gujarats/GenerateLocation.png)](http://travis-ci.org/Gujarats/GenerateLocation)
This Library is creating random location from given location. the result will be in decimal degree.
for example if you have a base location like `48.8588377,2.2775176`. You can generate new location using this library.

## How To Use It
```shell
$ go get github.com/Gujarats/GenerateLocation
```
### Example
let's say we wanted to generate new location based on the following latitude and longitude : `48.8588377,2.2775176` we can do this by using this code snippet : 
```go
package driver

import (
	"github.com/Gujarats/GenerateLocation"
	"github.com/icrowley/fake" // used for generate random names.
)

// We created Driver struct in this case for creating driver data with random names and location.
type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
}

// Return drivers data from given sum argument.
func GenereateDriver(sum int) []Driver {
	var drivers []Driver
	location.SetupLocation(48.8588377, 2.2775176)

	// get 30 % of the sum data
	smallPercentage := (30.0 / 100.0) * float64(sum)
	percentData := int(smallPercentage)

	// random lat lon based on seconds
	for i := 0; i <= sum; i++ {
		if sum-i <= percentData {
			// generate lat and lon using minute. from specific number 1-3
			lat, lon := location.RandomLatLongMinute(4)
			dummyDriver := Driver{
				Name:   fake.FullName(),
				Lat:    lat,
				Lon:    lon,
				Status: true,
			}
			drivers = append(drivers, dummyDriver)
		} else {
			// generate lat and lon using seconds. from specific number 1-6
			lat, lon := location.RandomLatLong(7)
			dummyDriver := Driver{
				Name:   fake.FullName(),
				Lat:    lat,
				Lon:    lon,
				Status: true,
			}
			drivers = append(drivers, dummyDriver)
		}

	}

	return drivers
}

```

#### Note
From the code above we called 
```go 
location.SetupLocation(Latitude,Longitude) 
``` 
to set base location. and then you can get the random location using

```go 
locattion.RandomLatLong(yourInputInSecondsRange) // get random lat lang using seconds; per second would be around 20 meters.
``` 
and 

```go 
locattion.RandomLatLongMinute(InputYourMinuteRange) // get random lat lang using seconds; per minute would be around 1 Km.
```

### Example 2
Another way to Genereate new latitude longitude : 
```go

// insert dummy location from latitude and longitude.
func insertDummyMarkLocation(cityName string, city *cityModel.City) {
	// some location in Bandung
	lat := -6.8647721
	lon := 107.553501
	var locations []location.Location

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (50/1)^2 = 2500 district
	newLocations = location.GenerateLocation(lat, lon, 1, 50)
}

```

## How does it work
First it will convert the decimal degree to degre like `48.8588377` to `48° 51' 23.8104''` and then adding the random value given to seconds and minute depends on what method you are using. After that It will convert back to decimal and you will get new latitude and longitude value.

## Contribution
I'm open to any improvement so please give your pull request if you have one :D.
### thanks to
* [Rolfl](http://codereview.stackexchange.com/a/156380/80799)
for his review.

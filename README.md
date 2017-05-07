# Generate Random Location [![Build Status](https://secure.travis-ci.org/Gujarats/GenerateLocation.png)](http://travis-ci.org/Gujarats/GenerateLocation) [![GoDoc](https://godoc.org/gopkg.in/gujarats/GenerateLocation?status.svg)](https://godoc.org/gopkg.in/gujarats/GenerateLocation.v1)
This Library is creating random location from given location. the result will be in decimal degree.
for example if you have a base location like `48.8588377,2.2775176`. You can generate new location using this library.

## How To Use It

First Download the repository
```shell
$ go get gopkg.in/gujarats/GenerateLocation.v1
```

And Import to your project 
```go
import "gopkg.in/gujarats/GenerateLocation.v1" // refers as location
```

### Example 1
Another way to Genereate new latitude longitude : 
```go

// insert dummy location from latitude and longitude.
func insertDummyMarkLocation(cityName string, city *cityModel.City) {
	// some location in Bandung
	lat := -6.8647721
	lon := 107.553501
    loc := location.New(lat,lon)

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (50/1)^2 = 2500 district
	newLocations = loc.GenerateLocation(1, 50)
}

```

### Example 2
let's say we wanted to generate new location based on the following latitude and longitude : `48.8588377,2.2775176` we can do this by using this code snippet : 
```go
package driver

import (
    "gopkg.in/gujarats/GenerateLocation.v1"
	"github.com/icrowley/fake" // used for generate random names.
)

// We created Driver struct in this case for creating driver data with random names and location.
type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
}

// Generate dummy drivers this will return []Driver with given sum.
// with new location from latitude and longitude given.
func GenereateDriver(lat, lon float64, sum int) []Driver {
	var drivers []Driver
	loc := location.New(lat, lon)

	// get 50 % of the sum data
	smallPercentage := (50.0 / 100.0) * float64(sum)
	percentData := int(smallPercentage)

	// random lat lon based on seconds
	for i := 0; i <= sum; i++ {
		if sum-i <= percentData {
			// generate lat and lon using minute. from specific number 1-3
			lat, lon := loc.RandomLatLongMinute(4)
			dummyDriver := Driver{
				Name:   fake.FullName(),
				Lat:    lat,
				Lon:    lon,
				Status: false,
			}
			drivers = append(drivers, dummyDriver)
		} else {
			// generate lat and lon using seconds. from specific number 1-6
			lat, lon := loc.RandomLatLongSeconds(7)
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

## Contribution
I'm open to any improvement so please give your pull request if you have one :D.
### thanks to
* [Rolfl](http://codereview.stackexchange.com/a/156380/80799)
for his review.

# Generate Random Location [![Build Status](https://secure.travis-ci.org/Gujarats/GenerateLocation.png)](http://travis-ci.org/Gujarats/GenerateLocation)
This Library is creating random location from given location. the result will be in decimal degree.
for example if you have a base location like 48.8588377,2.2775176. You can generate new location using this library.

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
	"github.com/icrowley/fake"
)

type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
}

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
``` first to set base location. and then you can get the random location using
```go 
locattion.RandomLatLong(yourInputInSecondsRange) 
``` 
and 

```go locattion.RandomLatLongMinute(InputYourMinuteRange) 
```

## How does it work
First it will convert the decimal degree to degre like `48° 51' 23.8104''` and then adding the random value given to seconds and minute depends on what method you are using. After that It will convert back to decimal and you will get new latitude and longitude value.

## Contribution
I'm open to any improvement os please give your pull request if you have one :D.

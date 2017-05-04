package location

import (
	"fmt"
	"testing"
)

func TestGetCenterLocation(t *testing.T) {
	testObjects := []struct {
		Lat                    float64
		Lon                    float64
		Distance               float64
		LimitLength            float64
		TotalExpectedLocations int
	}{
		// this should be 100 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, LimitLength: 50.0, TotalExpectedLocations: 100},

		// this should be 25 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2.0, LimitLength: 10.0, TotalExpectedLocations: 25},
	}

	for _, test := range testObjects {
		actualLocations := GetCenterLocation(test.Lat, test.Lon, test.Distance, test.LimitLength)

		fmt.Printf("result  = %+v\n", actualLocations)

	}

}

func TestGenerateLocation(t *testing.T) {
	testObjects := []struct {
		Lat                    float64
		Lon                    float64
		Distance               float64
		LimitLength            float64
		TotalExpectedLocations int
	}{
		// this should be 100 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, LimitLength: 50.0, TotalExpectedLocations: 100},

		// this should be 25 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2.0, LimitLength: 10.0, TotalExpectedLocations: 25},
	}

	for _, test := range testObjects {
		location := New(test.Lat, test.Lon)
		actualLocations := location.GenerateLocation(test.Distance, test.LimitLength)
		if len(actualLocations) != test.TotalExpectedLocations {
			//fmt.Printf("result  = %+v\n", actualLocations)
			t.Errorf("Error TotalExpectedLocations = %v, but actual = %v\n", test.TotalExpectedLocations, len(actualLocations))
		}
	}

}

func TestNewPoint(t *testing.T) {
	testObjects := []struct {
		Lat       float64
		Lon       float64
		Distance  float64
		Direction string
	}{
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, Direction: "west"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, Direction: "east"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, Direction: "north"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, Direction: "south"},
	}

	for _, test := range testObjects {
		lat, lon := newPoint(test.Lat, test.Lon, test.Distance, test.Direction)
		// point negative
		switch test.Direction {
		case "west":
			if !(lon < test.Lon) {
				t.Errorf("Error actual lon must be less. Actual lon= %v, current = %v", lon, test.Lon)
			}
		case "east":
			if !(lon > test.Lon) {
				t.Errorf("Error actual lon must be less. Actual lon= %v, current = %v", lon, test.Lon)
			}
		case "north":
			if !(lat > test.Lat) {
				t.Errorf("Error actual lat must be more. Actual Lat= %v, current Lat = %v", lat, test.Lat)
			}
		case "south":
			if !(lat < test.Lat) {
				t.Errorf("Error actual lat must be less. Actual Lat = %v, current Lat = %v", lat, test.Lat)
			}

		}

	}
}

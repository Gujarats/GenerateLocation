package location

import "testing"

func TestGenerateLocation(t *testing.T) {
	testObjects := []struct {
		Lat                    float64
		Lon                    float64
		Distance               int
		LimitLength            int
		TotalExpectedLocations int
	}{
		// this should be 100 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5, LimitLength: 50, TotalExpectedLocations: 100},

		// this should be 25 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2, LimitLength: 10, TotalExpectedLocations: 25},
	}

	for _, test := range testObjects {
		actualLocations := GenerateLocation(test.Lat, test.Lon, test.Distance, test.LimitLength)
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
		Distance  int
		Direction string
	}{
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5, Direction: "west"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5, Direction: "east"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5, Direction: "north"},
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5, Direction: "south"},
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

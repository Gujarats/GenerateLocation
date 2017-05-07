package location

import (
	"fmt"
	"math"
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

func TestGetMapLocations(t *testing.T) {
	testObjects := []struct {
		Lat         float64
		Lon         float64
		Distance    float64
		LimitLength float64
	}{
		// this should be 100 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, LimitLength: 50.0},

		// this should be 25 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2.0, LimitLength: 10.0},
	}

	for _, testObject := range testObjects {
		loc := New(testObject.Lat, testObject.Lon)
		mapLocations := loc.GenerateLocation(testObject.Distance, testObject.LimitLength)

		ExpectedLength := int(math.Pow((testObject.LimitLength / testObject.Distance), 2))
		if len(mapLocations) != ExpectedLength {
			t.Errorf("Error length location = %+v , ExpectedLegth = %+v\n", len(mapLocations), ExpectedLength)
		}

	}

}

func TestgetCenterLocation(t *testing.T) {
	testObjects := []struct {
		MapLocations          [][]Location
		CenterIndex           [2]int
		Position              int
		ExpectedMarkedCenters [4][2]int
		ExpectedLocations     [4]CenterLocation
	}{
		{
			MapLocations: [][]Location{
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}},
			},
			CenterIndex: [2]int{2, 2},
			Position:    1,
			ExpectedMarkedCenters: [4][2]int{
				{1, 3},
				{1, 1},
				{3, 1},
				{3, 3},
			},
			ExpectedLocations: [4]CenterLocation{
				{
					Quadran:        1,
					MarkedLocation: Location{0.7, 0.8},
				},
				{
					Quadran:        2,
					MarkedLocation: Location{0.3, 0.4},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.3, 0.4},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.7, 0.8},
				},
			},
		},
		{
			MapLocations: [][]Location{
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}},
			},
			CenterIndex: [2]int{3, 3},
			Position:    1,
			ExpectedMarkedCenters: [4][2]int{
				{2, 4},
				{2, 2},
				{4, 2},
				{4, 4},
			},
			ExpectedLocations: [4]CenterLocation{
				{
					Quadran:        1,
					MarkedLocation: Location{0.9, 0.10},
				},
				{
					Quadran:        2,
					MarkedLocation: Location{0.5, 0.6},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.5, 0.6},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.9, 0.10},
				},
			},
		},
		{
			MapLocations: [][]Location{
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
				{{0.1, 0.2}, {0.3, 0.4}, {0.5, 0.6}, {0.7, 0.8}, {0.9, 0.10}, {0.11, 0.12}, {0.13, 0.14}, {0.15, 0.16}},
			},
			CenterIndex: [2]int{4, 4},
			Position:    2,
			ExpectedMarkedCenters: [4][2]int{
				{2, 6},
				{2, 2},
				{6, 2},
				{6, 6},
			},
			ExpectedLocations: [4]CenterLocation{
				{
					Quadran:        1,
					MarkedLocation: Location{0.13, 0.14},
				},
				{
					Quadran:        2,
					MarkedLocation: Location{0.5, 0.6},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.5, 0.6},
				},
				{
					Quadran:        3,
					MarkedLocation: Location{0.13, 0.14},
				},
			},
		},
	}

	for _, testObject := range testObjects {
		markedCenters, locations := getCenterLocations(testObject.MapLocations, testObject.CenterIndex, testObject.Position)

		// check the length it must be the same lenght
		if len(locations) != len(testObject.ExpectedLocations) {
			t.Errorf("Error length should be the same !. result = %+v , expected = %+v\n", len(locations), len(testObject.ExpectedLocations))
		}

		// check MarkedLocations
		for index, markedLocation := range markedCenters {
			if markedLocation != testObject.ExpectedMarkedCenters[index] {
				t.Errorf("Error marked location = %+v, expected = %+v\n", markedLocation, testObject.ExpectedLocations[index])
			}
		}

		// check locations results it must be equal to the ExpectedLocations
		for index, location := range locations {
			// comparing struct from the actual result and expected
			if location != testObject.ExpectedLocations[index] {
				t.Errorf("Error location = %+v, expected = %+v\n", location, testObject.ExpectedLocations[index])
			}
		}
	}
}

func TestGetCenterQuadranLocations(t *testing.T) {
	testObjects := []struct {
		Lat         float64
		Lon         float64
		Distance    float64
		LimitLength float64
		DeepLevel   int
	}{
		// this should be 100 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 5.0, LimitLength: 50.0, DeepLevel: 2},

		// this should be 25 locations
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2.0, LimitLength: 10.0, DeepLevel: 1},

		// index 2 gives error
		{Lat: -6.9875393, Lon: 108.4446289, Distance: 2.0, LimitLength: 10.0, DeepLevel: 10},
	}

	for index, testObject := range testObjects {
		location := New(testObject.Lat, testObject.Lon)
		mapLocations, err := location.GetCenterQuadranLocations(testObject.Distance, testObject.LimitLength, testObject.DeepLevel)

		// expecting errro from this case.
		if index == 2 {
			if err == nil {
				t.Errorf("Expected got = %+v\n", err)
			}
			return
		}

		// check the every mapLocations to DeepLevel
		for level := 0; level < testObject.DeepLevel; level++ {
			// we expecte the result is not nil
			if mapLocations[level] == [4]CenterLocation{} {
				t.Errorf("Error expecting result from the index = %+v, result = %+v\n", index, mapLocations[level])
			}
		}
	}

}

// NOTE :  negative here is define the south and west here
// south and east are positive.
// we give the dummy(not real latitude and longitude) location here to speed up the writting test, but it works just exactly like real lat and lon.
func TestGetQuadranPosition(t *testing.T) {
	testObjects := []struct {
		BaseLocation    Location
		InputLocation   Location
		QuadranPosition string
	}{
		// index 0
		{
			BaseLocation:    Location{0.0, 0.0},
			InputLocation:   Location{0.1, 0.2},
			QuadranPosition: "q1",
		},

		// index 1
		{
			BaseLocation:    Location{-2.2, 2.0},
			InputLocation:   Location{0.1, 2.2},
			QuadranPosition: "q1",
		},

		// index 2
		{
			BaseLocation:    Location{-3.0, -7.0},
			InputLocation:   Location{0.1, 0.2},
			QuadranPosition: "q1",
		},

		// index 3
		{
			BaseLocation:    Location{3.0, -4.0},
			InputLocation:   Location{4.1, 4.2},
			QuadranPosition: "q1",
		},

		// index 4
		{
			BaseLocation:    Location{3.0, -4.0},
			InputLocation:   Location{-4.1, 4.2},
			QuadranPosition: "q2",
		},

		// index 5
		{
			BaseLocation:    Location{3.0, -4.0},
			InputLocation:   Location{2.9, 4.2},
			QuadranPosition: "q2",
		},
		// index 6
		{
			BaseLocation:    Location{3.0, -4.0},
			InputLocation:   Location{2.9, -4.2},
			QuadranPosition: "q3",
		},
		// index 7
		{
			BaseLocation:    Location{3.0, -4.0},
			InputLocation:   Location{3.9, -4.2},
			QuadranPosition: "q4",
		},
	}

	for index, testObject := range testObjects {
		quadranPosition, err := GetQuadranPosition(testObject.BaseLocation, testObject.InputLocation)
		if err != nil {
			t.Errorf("Error = %v , at index = %v\n", err.Error(), index)
		}

		if quadranPosition != testObject.QuadranPosition {
			t.Errorf("Error at index =%v ,quadran = %v, expected = %v\n", index, quadranPosition, testObject.QuadranPosition)
		}
	}
}

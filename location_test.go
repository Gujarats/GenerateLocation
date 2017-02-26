package location

import "testing"

func init() {
	SetupLocation(48.8588377, 2.2775176)
}

func TestToDegree(t *testing.T) {
	type testObject struct {
		InputLat float64
		Expected LocationDegree
	}

	testObjects := []testObject{
		{InputLat: 48.76,
			Expected: LocationDegree{
				Degree:  48,
				Minutes: 45,
				Seconds: 36,
			},
		},
	}

	for _, test := range testObjects {
		actual := ToDegree(test.InputLat)
		if actual.Degree != test.Expected.Degree || actual.Minutes != test.Expected.Minutes || actual.Seconds != test.Expected.Seconds {
			t.Errorf("Error actual = %+v, expected = %+v\n", actual, test.Expected)
		}
	}
}

func TestToDecimal(t *testing.T) {
	type testObject struct {
		InputLocation LocationDegree
		Expected      float64
	}

	testObjects := []testObject{
		{InputLocation: LocationDegree{48, 45, 36}, Expected: 48.76},
	}

	for _, test := range testObjects {
		actual := ToDecimal(test.InputLocation)

		if !isFloatEqual(actual, test.Expected) {
			t.Errorf("Error result is not the same actual = %v, expected = %v\n", actual, test.Expected)
		}
	}

}

func TestRandomLatLong(t *testing.T) {
	// init location paris
	rLat, rLon := RandomLatLong(12) // give 12 to get random number from 1-12

	// comparing the lat with rLat (generated value location)
	if isFloatEqual(rLat, lat) && isFloatEqual(rLon, lon) {
		t.Errorf("Error the generated value is the same")
	}

}

func TestRandonLatLongMinute(t *testing.T) {
	rLat, rLon := RandomLatLongMinute(6) // give 12 to get random number from 1-12

	// comparing the lat with rLat (generated value location)
	if isFloatEqual(rLat, lat) && isFloatEqual(rLon, lon) {
		t.Errorf("Error the generated value is the same")
	}

}

// We use EPILOPS to determice the precission of float number.
func isFloatEqual(a, b float64) bool {
	EPILOPS := 0.000001

	if (a-b) > EPILOPS || (b-a) > EPILOPS {
		return false
	}

	return true

}

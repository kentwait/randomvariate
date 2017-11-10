package randomvariate

import "math"

// Round rounds a value from zero based on the given units place.
// For example, the value to be round is 0.05 and the given unit is 0.1,
// Round will return 0.1 as it rounds 0.05 to the nearest tenths place.
func Round(x float64, e int) float64 {
	unit := math.Pow10(e)
	v := float64(int64(math.Abs(x)/unit+0.5)) * unit
	if math.Signbit(x) {
		return v * -1
	}
	return v
}

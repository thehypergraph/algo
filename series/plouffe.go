package series

import "math"

func Plouffe(max int) float64 {
	var pi float64

	for i := range max {
		pi += 1.0 / math.Pow(16, float64(i)) * (4.0/(8*float64(i)+1) - 2.0/(8*float64(i)+4) - 1.0/(8*float64(i)+5) - 1.0/(8*float64(i)+6))
	}

	return pi
}

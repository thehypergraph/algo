package geometry

import "math"

func L2Norm(nums []float64) float64 {
	sum := 0.0

	for x := range nums {
		sum = math.Pow(float64(x), 2)
	}

	return math.Sqrt(sum)
}

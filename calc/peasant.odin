package calc

import "core:testing"

// Russian peasant multiplication
peasant_multiply :: proc(a, b: int) -> int {
	a, b := a, b // shadow params to make them mutable
	result := 0
	for a > 0 {
		if a % 2 == 1 {
			result += b
		}
		a /= 2
		b *= 2
	}
	return result
}

@(test)
peasant_multiply_test :: proc(t: ^testing.T) {
	expect := 4 * 5

	testing.expect(t, peasant_multiply(4, 5) == expect, "peasant_multiply(4, 5) should equal 20")
}

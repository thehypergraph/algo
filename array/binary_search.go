package array

import "errors"

func BinarySearch(haystack []uint32, needle uint32) (uint32, error) {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := haystack[m]

		if v == needle {
			return v, nil
		}

		if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return 0, errors.New("No matches")
}

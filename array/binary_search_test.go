package array

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// Test with sorted array - element found
	arr := []uint32{1, 3, 5, 7, 9, 11, 13, 15}

	// Test finding elements at different positions
	tests := []struct {
		target   uint32
		expected uint32
		hasError bool
	}{
		{1, 1, false},   // first element
		{15, 15, false}, // last element
		{7, 7, false},   // middle element
		{5, 5, false},   // early element
		{13, 13, false}, // late element
	}

	for _, test := range tests {
		result, err := BinarySearch(arr, test.target)
		if test.hasError {
			if err == nil {
				t.Errorf("BinarySearch(%v, %d) expected error but got none", arr, test.target)
			}
		} else {
			if err != nil {
				t.Errorf("BinarySearch(%v, %d) unexpected error: %v", arr, test.target, err)
			}
			if result != test.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
			}
		}
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []uint32{1, 3, 5, 7, 9, 11, 13, 15}

	// Test elements not in array
	notFoundTests := []uint32{0, 2, 4, 6, 8, 10, 12, 14, 16, 100}

	for _, target := range notFoundTests {
		result, err := BinarySearch(arr, target)
		if err == nil {
			t.Errorf("BinarySearch(%v, %d) = %d, nil; expected error", arr, target, result)
		}
		if result != 0 {
			t.Errorf("BinarySearch(%v, %d) = %d; expected 0 when not found", arr, target, result)
		}
	}
}

func TestBinarySearchEdgeCases(t *testing.T) {
	// Test empty array
	emptyArr := []uint32{}
	result, err := BinarySearch(emptyArr, 5)
	if err == nil {
		t.Errorf("BinarySearch([], 5) expected error but got none")
	}
	if result != 0 {
		t.Errorf("BinarySearch([], 5) = %d; expected 0", result)
	}

	// Test single element array - found
	singleArr := []uint32{42}
	result, err = BinarySearch(singleArr, 42)
	if err != nil {
		t.Errorf("BinarySearch([42], 42) unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("BinarySearch([42], 42) = %d; expected 42", result)
	}

	// Test single element array - not found
	result, err = BinarySearch(singleArr, 5)
	if err == nil {
		t.Errorf("BinarySearch([42], 5) expected error but got none")
	}
	if result != 0 {
		t.Errorf("BinarySearch([42], 5) = %d; expected 0", result)
	}

	// Test two element array
	twoArr := []uint32{10, 20}
	result, err = BinarySearch(twoArr, 10)
	if err != nil {
		t.Errorf("BinarySearch([10, 20], 10) unexpected error: %v", err)
	}
	if result != 10 {
		t.Errorf("BinarySearch([10, 20], 10) = %d; expected 10", result)
	}

	result, err = BinarySearch(twoArr, 20)
	if err != nil {
		t.Errorf("BinarySearch([10, 20], 20) unexpected error: %v", err)
	}
	if result != 20 {
		t.Errorf("BinarySearch([10, 20], 20) = %d; expected 20", result)
	}

	result, err = BinarySearch(twoArr, 15)
	if err == nil {
		t.Errorf("BinarySearch([10, 20], 15) expected error but got none")
	}
}

func TestBinarySearchDuplicates(t *testing.T) {
	// Test array with duplicate elements
	arr := []uint32{1, 2, 2, 2, 3, 4, 5}

	// Should return the value when found
	result, err := BinarySearch(arr, 2)
	if err != nil {
		t.Errorf("BinarySearch(%v, 2) unexpected error: %v", arr, err)
	}
	if result != 2 {
		t.Errorf("BinarySearch(%v, 2) = %d; expected 2", arr, result)
	}
}

func TestBinarySearchLargeNumbers(t *testing.T) {
	// Test with larger uint32 numbers
	arr := []uint32{100, 500, 1000, 5000, 10000, 50000, 100000}

	tests := []struct {
		target   uint32
		expected uint32
		hasError bool
	}{
		{100, 100, false},
		{1000, 1000, false},
		{100000, 100000, false},
		{999, 0, true},    // not found
		{200000, 0, true}, // not found
	}

	for _, test := range tests {
		result, err := BinarySearch(arr, test.target)
		if test.hasError {
			if err == nil {
				t.Errorf("BinarySearch(%v, %d) expected error but got none", arr, test.target)
			}
		} else {
			if err != nil {
				t.Errorf("BinarySearch(%v, %d) unexpected error: %v", arr, test.target, err)
			}
			if result != test.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
			}
		}
	}
}

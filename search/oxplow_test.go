package search

import "testing"

func TestOxPlow(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	target := 5
	expected := true
	result, _ := OxPlow(data, target)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestOxPlowNotFound(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	target := 10
	expected := false
	result, _ := OxPlow(data, target)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestOxPlowUnevenRows(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6, 8},
		{9, 10, 11},
	}
	target := 8
	expected := true
	result, _ := OxPlow(data, target)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

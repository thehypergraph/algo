package geometry

import (
	"testing"
)

func TestL2Norm(t *testing.T) {
	nums := []float64{1.0, 2.0, 3.0}
	result := L2Norm(nums)
	expected := 2.0

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

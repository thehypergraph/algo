package series

import (
	"testing"
)

func TestFib(t *testing.T) {

	result := Fib(10)

	if result != 55 {
		t.Errorf("Expected Fib(10) to return 55, got %v", result)
	}
}

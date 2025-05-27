package series

import (
	"testing"
)

func TestPrimes(t *testing.T) {

	result := NthPrime(10)

	if result != 29 {
		t.Errorf("Expected NthPrime(10) to return 29, got %v", result)
	}

	result = NthPrime(5)

	if result != 11 {
		t.Errorf("Expected NthPrime(5) to return 11, got %v", result)
	}

	result = NthPrime(1)

	if result != 2 {
		t.Errorf("Expected NthPrime(1) to return 2, got %v", result)
	}
}

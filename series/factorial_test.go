package series

import "testing"

func TestFactorial(t *testing.T) {
	got := Factorial(5)
	want := 120

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

package series

import "testing"

func TestFib(t *testing.T) {
	got := Fib(5)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

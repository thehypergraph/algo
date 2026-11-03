package series

import "core:strings"
import "core:testing"

fib :: proc(n: int) -> int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}

@(test)
fib_test :: proc(t: ^testing.T) {
	expect := 55
	result := fib(10)
	testing.expect(t, result == expect)
}

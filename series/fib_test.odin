package series

import "core:testing"
import "core:fmt"

@(test)
test_fib :: proc(t: ^testing.T) {
	res := fib(10)
	fmt.println(res)
	testing.expect(t, res == 55)
}

package series

var cache = map[int32]int32{
	0: 1,
	1: 1,
}

func Fib(n int32) int32 {
	item, exists := cache[n]

	if exists {
		return item
	}

	return Fib(n-1) + Fib(n-2)
}

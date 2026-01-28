package series

import (
	"fmt"
	"math"
)

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

func BinetFib(n int32) int32 {

	a := 1 / math.Sqrt(5)
	b := math.Pow((1+math.Sqrt(5))/2, float64(n))
	c := math.Pow((1-math.Sqrt(5))/2, float64(n))

	res := a * (b - c)

	fmt.Println(res)

	return int32(res)
}

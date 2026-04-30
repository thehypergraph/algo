package main

import (
	"algo/series"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	for i := int32(1); i <= 20; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, series.Fib(i))
	}
}

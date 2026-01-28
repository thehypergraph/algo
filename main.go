package main

import (
	"algo/series"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	for i := 1; i <= 100; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, series.BinetFib(int32(i)))
	}
}

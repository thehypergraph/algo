package main

import (
	series "algo/series"
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		tmp := rand.Intn(100)
		prime := series.NthPrime(tmp)
		fmt.Println(prime)
	}

}

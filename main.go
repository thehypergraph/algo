package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	tmp := rand.Intn(100)
	// 	prime := series.NthPrime(tmp)
	// 	fmt.Println(prime)
	// }

	vals := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}

	for i := 0; i < len(vals); i++ {
		for j := 0; j < len(vals[i]); j++ {
			fmt.Println("Index: ", i, j, "Value: ", vals[i][j])
		}
	}
}

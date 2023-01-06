package algo

func Factorial(n uint32) uint32 {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

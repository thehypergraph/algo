package series

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// NthPrime returns the nth prime number
func NthPrime(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	num := 2

	for {
		if IsPrime(num) {
			count++
			if count == n {
				return num
			}
		}
		num++
	}
}

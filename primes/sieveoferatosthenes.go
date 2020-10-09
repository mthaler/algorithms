package primes

import "math"

func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, 0)
	for i := 2; i <= n; i++ {
		primes = append(primes, true)
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i <= max; i++ {
		if primes[i - 2] {
			for j := i * i; j <= n; j += i {
				primes[j - 2] = false
			}
		}
	}
	result := make([]int, 0)
	for i := 2; i <= n; i++ {
		if primes[i - 2] {
			result = append(result, i)
		}
	}
	return result
}

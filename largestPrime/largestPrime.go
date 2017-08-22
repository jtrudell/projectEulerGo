// projecteuler.net Problem 3: Largest prime factor
// What is the largest prime factor of the number 600851475143?

package main

import (
	"fmt"
	"math"
)

func main() {
	num := 600851475143
	primes := makeSieve(num)
	factors := primeFactors(primes, num)

	fmt.Println("largest prime of", num, "is", factors[len(factors)-1])
}

// Generate sieve of Eratosthenes to find all primes up to square root of 600851475143
func makeSieve(n int) []int {
	n = int(math.Sqrt(float64(n)))
	var sieve = make([]bool, n)
	var primes = make([]int, 1)

	for i, value := range sieve {
		if i < 2 {
			continue
		}

		// If value is false it is a prime, append index (which represents the prime number) to primes.
		// Iterate through indexes divisible by the prime, setting their value to true as they are not prime.
		if value == false {
			primes = append(primes, i)
			for j := i * i; j < n; j += i {
				sieve[j] = true
			}
		}

	}
	return primes
}

// Find prime factors of 600851475143
func primeFactors(primes []int, n int) []int {
	var factors = make([]int, 1)

	for _, val := range primes {
		if val != 0 && n%val == 0 {
			factors = append(factors, val)
		}
	}
	return factors
}

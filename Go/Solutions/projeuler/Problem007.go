package euler

import (
	"fmt"
	"time"
)

/*
Problem007 answers the problem at : https://projecteuler.net/problem=7

* Problem 7:
		By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
		What is the 10001st prime number?

*/
func Problem007() {

	start := time.Now()
	ans := PrimeSieve(1000000)[10000]
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 7 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

//PrimeSieve implements the sieve of Eratosthenes
func PrimeSieve(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] == true {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}
	return primes
}

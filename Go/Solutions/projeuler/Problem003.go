package euler

import (
	"fmt"
	"time"
)

/*
Problem003 answers the problem at : https://projecteuler.net/problem=3

* Problem 3:
		The prime factors of 13195 are 5, 7, 13 and 29.
		What is the largest prime factor of the number 600851475143 ?
*/
func Problem003() {
	start := time.Now()
	ans := _maxPrime(Factors(600851475143))
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 3 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

package euler

import (
	"fmt"
	"time"
)

/*
Problem010 answers the problem at : https://projecteuler.net/problem=10

* Problem 10:
		Find the sum of all the primes below two million.
*/
func Problem010() {

	start := time.Now()
	ans := Sum(PrimeSieve(2000000))
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 10 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


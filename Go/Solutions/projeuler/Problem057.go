package euler

import (
	"fmt"
	"math"
	"time"
)

/*
Problem057 answers the problem at : https://projecteuler.net/problem=57

* Problem 57:

It is possible to show that the square root of two can be expressed as an infinite continued fraction.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than the denominator?
*/
func Problem057() {
	start := time.Now()
	n, a, ans := 3, 2, 0

	for i := 2; i < 1001; i++ {
		a = n + a
		n = n + 2*a

		if math.Log2(float64(n)) > math.Log2(float64(a)) { // needs work -_-
			ans++
		}

	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 57 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

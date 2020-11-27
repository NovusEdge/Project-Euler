package euler

import (
	"fmt"
	"math"
	"time"
)

var _amicableNums = []int{}

/*
Problem021 answers the problem at : https://projecteuler.net/problem=21

* Problem 21:
		Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
		If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

		For example, d(220) = 284 and d(284) = 220, so 220 and 284 form an amicable pair.

		Evaluate the sum of all the amicable numbers under 10000.
*/
func Problem021() {
	start := time.Now()
	ans := 0

	for i := 1; i < 10000; i++ {
		for j := 1; j <= i; j++ {
			if isAmicablePair(i, j) {
				ans += (i + j)
			}
		}
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 21 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func isAmicablePair(a, b int) bool {
	return d(a) == b && d(b) == a && a != b
}

func d(n int) (res int) {
	upper := int(math.Sqrt(float64(n)))
	for i := 1; i <= upper; i++ {
		if n%i == 0 {
			res += i + (n / i)
		}
	}
	res -= n
	return
}

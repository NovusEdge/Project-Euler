package euler

import (
	"fmt"
	"time"
)

/*
Problem034 answers the problem at : https://projecteuler.net/problem=34

* Problem 34:

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
*/
func Problem034() {
	start := time.Now()

	ans := 0
	for i := 3; i < 100000; i++ {
		if check34(i) {
			ans += i
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 34 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

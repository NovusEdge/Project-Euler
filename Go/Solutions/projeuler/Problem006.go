package euler

import (
	"fmt"
	"time"
)

/*
Problem006 answers the problem at : https://projecteuler.net/problem=6

* Problem 6:
		Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func Problem006() {

	start := time.Now()
	ans := sumSquare(100) - squareSum(100)
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 6 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


package euler

import (
	"fmt"
	"time"
)

/*
Problem052 answers the problem at : https://projecteuler.net/problem=52

* Problem 52:

	It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

	Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/
func Problem052() {
	start := time.Now()
	i := 100000

	for !(sameDigits(i, 2*i) && sameDigits(i, 3*i) && sameDigits(i, 4*i) && sameDigits(i, 5*i) && sameDigits(i, 6*i)) {
		i++
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 52 : %d\n", i)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

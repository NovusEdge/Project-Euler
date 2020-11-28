package euler

import (
	"fmt"
	"time"
)

/*
Problem005 answers the problem at : https://projecteuler.net/problem=5

* Problem 5:
		2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

		What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func Problem005() {

	start := time.Now()

	ans := LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 5 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

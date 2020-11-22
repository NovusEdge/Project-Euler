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

	var ans int
	start := time.Now()
	getAns5(&ans)
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 5 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


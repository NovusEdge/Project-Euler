package euler

import (
	"fmt"
	"time"
)

/*
Problem014 answers the problem at : https://projecteuler.net/problem=14

* Problem 14:

	The following iterative sequence is defined for the set of positive integers:

	n → n/2 (n is even)
	n → 3n + 1 (n is odd)

	Using the rule above and starting with 13, we generate the following sequence:
			13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

	It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

	Which starting number, under one million, produces the longest chain?

* NOTE: Once the chain starts the terms are allowed to go above one million.
*/
func Problem014() {
	start := time.Now()
	ans := getMaxFromSeq(getAllSeq())
	end := time.Now()
	fmt.Printf("\nAnswer to Problem 14 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


package euler

import (
	"fmt"
	"sort"
	"time"
)

/*
Problem024 answers the problem at : https://projecteuler.net/problem=24

* Problem 24:
		If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
		The lexicographic permutations of 0, 1 and 2 are:

						012   021   102   120   201   210

		Find the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9.
*/
func Problem024() {

	start := time.Now()
	perms := getPerms("0123456789")
	sort.Strings(perms)

	ans := perms[999999]
	end := time.Now()
	fmt.Printf("\nAnswer to Problem 24 : %s\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


package euler

import (
	"fmt"
	"time"
)

/*
Problem009 answers the problem at : https://projecteuler.net/problem=9

* Problem 9:
		There exists exactly one Pythagorean triplet for which a + b + c = 1000 (a < b < c).
		Find the product abc.
*/
func Problem009() {

	start := time.Now()
	ans := getAns9()
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 9 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

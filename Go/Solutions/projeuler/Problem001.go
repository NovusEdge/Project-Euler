package euler

import (
	"fmt"
	"time"
)

/*
Problem001 answers the problem at : https://projecteuler.net/problem=1

* Problem 1:
		If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
		Find the sum of all the multiples of 3 or 5 below 1000.
*/
func Problem001() {

	start := time.Now()
	ans := 0
	for i := 2; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			ans += i
		}
	}
	end := time.Now()
	fmt.Printf("\nAnswer to Problem 1 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

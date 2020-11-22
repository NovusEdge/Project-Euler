package euler

import (
	"fmt"
	"time"
)

/*
Problem004 answers the problem at : https://projecteuler.net/problem=4

* Problem 4:
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

		Find the largest palindrome made from the product of two 3-digit numbers.
*/
func Problem004() {
	start := time.Now()
	tempArr := []int{}

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			if IsPall(i * j) {
				tempArr = append(tempArr, i*j)
			}
		}
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 4 : %d\n", Max(tempArr))
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

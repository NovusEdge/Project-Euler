package euler

import (
	"fmt"
	"time"
)

/*
Problem033 answers the problem at : https://projecteuler.net/problem=33

* Problem 33:

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

***We shall consider fractions like, 30/50 = 3/5, to be trivial examples.***

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.

*/
func Problem033() {
	res := [][]int{}

	start := time.Now()

	for a := 11; a < 100; a++ {
		for b := 11; b < 100; b++ {
			if check33(a, b) {
				a1, b1, _ := cut(a, b)
				res = append(res, []int{a1, b1})
			}
		}
	}

	end := time.Now()
	fmt.Println("\nAnswer to Problem 33 (Resultant fractions) : ", res)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

package euler

import (
	"fmt"
	"math/big"
	"time"
)

/*
Problem053 answers the problem at : https://projecteuler.net/problem=53

* Problem 53:

There are exactly ten ways of selecting three from five, 12345:
	123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

How many, not necessarily distinct, values of C(n, r) for, 1 ≤ n ≤ 100 are greater than one-million?
*/
func Problem053() {
	start := time.Now()
	ans := 0

	for a := 1; a <= 100; a++ {
		for r := 1; r <= a; r++ {
			if C(a, r).Cmp(big.NewInt(1000000)) > 0 {
				ans++
			}
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 53 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


package euler

import (
	"fmt"
	"math/big"
	"time"
)

/*
Problem056 answers the problem at : https://projecteuler.net/problem=56

* Problem 56:

Considering natural numbers of the form, a**b, where a, b < 100, what is the maximum digital sum?
*/
func Problem056() {
	start := time.Now()
	ans := int64(0)

	for i := 99; i >= 1; i-- {
		for j := 99; j >= 1; j-- {
			k := big.NewInt(0)
			k.Exp(big.NewInt(int64(i)), big.NewInt(int64(j)), nil)
			temp := DigitalRoot(k)
			if temp > ans {
				ans = temp
			}
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 56 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

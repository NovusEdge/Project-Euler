package euler

import (
	"fmt"
	"strconv"
	"time"
)

/*
Problem036 answers the problem at : https://projecteuler.net/problem=36

* Problem 36:

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

**(Please note that the palindromic number, in either base, may not include leading zeros.)**

*/
func Problem036() {
	start := time.Now()
	ans := 0

	for i := 1; i < 1000000; i++ {
		k := strconv.Itoa(i)
		if isPallStr(k) && isPallStr(bin(int64(i))) {
			ans += i
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 36 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


package euler

import (
	"fmt"
	"time"
)

/*
Problem027 answers the problem at : https://projecteuler.net/problem=27

* Problem 27:
Considering quadratics of the form:
		n^2 + a*n + b , where |a| < 1000 and |b| <= 1000

Find the product of the coefficients, *a*
and *b*, for the quadratic expression that produces the maximum number of primes for consecutive values of *n*, starting with *n* = 0.
*/
func Problem027() {
	quadDict := map[int]string{}
	start := time.Now()

	for a := -1000; a < 1001; a++ {
		for b := -1000; b < 1001; b++ {
			c, q := formQuad(a, b)
			if q == "" {
				continue
			} else {
				quadDict[c] = q
			}
		}
	}

	ans := quadDict[dictMax(quadDict)]
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 27 : %s\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func formQuad(a, b int) (int, string) {
	counter, i := 0, 0
	for true {
		exprn := i*i + a*i + b
		if IsPrime(int64(exprn)) {
			counter++
			i++
		} else {
			return counter, fmt.Sprintf("n**2 + %dn + %d", a, b)
		}
	}
	return 0, ""
}

func dictMax(m map[int]string) int {
	flag := 0
	for k := range m {
		if k > flag {
			flag = k
		}
	}
	return flag
}

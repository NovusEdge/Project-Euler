package euler

import (
	"fmt"
	"strconv"
	"time"
)

/*
Problem037 answers the problem at : https://projecteuler.net/problem=37

* Problem 37:

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/
func Problem037() {
	start := time.Now()
	ans, i := int64(0), int64(10)
	counter := 0

	for counter < 11 {
		if isTruncR(i) && isTruncL(i) {
			ans += i
			counter++
		}
		i++
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 37 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func isTruncL(n int64) bool {
	s := []byte(strconv.FormatInt(n, 10))

	for i := 0; i < len(s); i++ {
		k, _ := strconv.Atoi(string(s[i:]))
		if !IsPrime(int64(k)) {
			return false
		}
	}
	return true
}

func isTruncR(n int64) bool {
	s := []byte(strconv.FormatInt(n, 10))

	for i := len(s); i > 0; i-- {
		k, _ := strconv.Atoi(string(s[:i]))
		if !IsPrime(int64(k)) {
			return false
		}
	}
	return true
}

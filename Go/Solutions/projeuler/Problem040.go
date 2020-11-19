package euler

import (
	"fmt"
	"time"
)

/*
Problem040 answers the problem at : https://projecteuler.net/problem=40

* Problem 40:

An irrational decimal fraction is created by concatenating the positive integers:

		0.123456789101112131415161718192021 ...

It can be seen that the 12th digit of the fractional part is 1.

If d(n) represents the nth digit of the fractional part, find the value of the following expression.

d(1) × d(10) × d(100) × d(1000) × d(10000) × d(100000) × d(1000000)
*/
func Problem040() {
	start := time.Now()
	c := []int{}

	for i := 0; i < 200000; i++ {
		c = append(c, Digits(i)...)
	}

	d40 := func(n int) int {
		return c[n]
	}

	ans := d40(1) * d40(10) * d40(100) * d40(1000) * d40(10000) * d40(100000) * d40(1000000)

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 40 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

package euler

import (
	"fmt"
	"time"
)

/*
Problem025 answers the problem at : https://projecteuler.net/problem=25

* Problem 25:
The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
func Problem025() {
	start := time.Now()
	ans := 0
	for true {
		if numLen(Fibonacci(int64(ans))) == 1000 {
			break
		}
		ans++
	}
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 25 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

package euler

import (
	"fmt"
	"time"
)

/*
Problem041 answers the problem at : https://projecteuler.net/problem=41

* Problem 41:

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

***What is the largest n-digit pandigital prime that exists?***
*/
func Problem041() {
	start := time.Now()

	ans := 0
	// fmt.Println(formNum([]int{1, 2, 2, 9, 5, 6, 7}))

	k := []int{12345, 123456, 1234567, 12345678, 123456789}
	for _, i := range k {
		p := getPrimePerm(i)
		if ans < p {
			ans = p
		}
	}
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 41 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

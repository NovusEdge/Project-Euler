package euler

import (
	"fmt"
	"time"
)

/*
Problem030 answers the problem at : https://projecteuler.net/problem=30

* Problem 30:

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.
The sum of these numbers is: 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/
func Problem030() {
	nums := []int{}
	start := time.Now()
	for i := 1000; i < 999999; i++ {
		if i == fifthPowSum(Digits(i)) {
			nums = append(nums, i)
		}
	}

	ans := Sum(nums)
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 30 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

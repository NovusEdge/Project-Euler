package euler

import (
	"fmt"
	"time"
)

/*
Problem023 answers the problem at : https://projecteuler.net/problem=23
	* Problem 23:

		A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

		A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

		As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

		Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/
func Problem023() {
	start := time.Now()
	abundantNums := getAbundant()
	boolMap := make(map[int]bool)

	for i := 1; i < 28123; i++ {
		boolMap[i] = true
	}

	ans := 0

	for _, i := range abundantNums {
		for _, j := range abundantNums {
			boolMap[i+j] = false
		}
	}

	for k, v := range boolMap {
		if v {
			ans += k
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 23 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

func isAbundant(n int) bool {
	var res int64
	k := Factors(int64(n))
	for _, i := range k {
		res += i
	}
	return res-int64(n) > int64(n)
}

func getAbundant() (res []int) {
	for i := 1; i < 28124; i++ {
		if isAbundant(i) {
			res = append(res, i)
		}
	}
	return
}

package euler

import (
	"fmt"
	"math"
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

func getPrimePerm(n int) (res int) {
	perms := permutations(Digits(n))
	for _, i := range perms {
		num := formNum(i)
		if IsPrime(int64(num)) && res < num {
			res = num
		}
	}
	return
}

func formNum(digits []int) (res int) {
	l := len(digits)
	res += digits[0] * int(math.Pow(10, float64(l)))

	for i := 1; i < l; i++ {
		p := int(math.Pow(10, float64(l-i)))
		res += digits[i] * p
	}
	res = res / 10
	return
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

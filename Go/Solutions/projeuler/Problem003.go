package euler

import (
	"fmt"
	"math"
	"time"
)

/*
Problem003 answers the problem at : https://projecteuler.net/problem=3

* Problem 3:
		The prime factors of 13195 are 5, 7, 13 and 29.
		What is the largest prime factor of the number 600851475143 ?
*/
func Problem003() {
	start := time.Now()
	ans := _maxPrime(Factors(600851475143))
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 3 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

//Factors returns an array of the factors of a number
func Factors(n int64) (res []int64) {
	upper := int64((math.Sqrt(math.Abs(float64(n)))))
	for i := int64(1); i < upper; i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, int64(n/i))
		}
	}
	return
}

//IsPrime reports if a number is a prime number or not
func IsPrime(n int64) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}
	upper := int64(math.Trunc(math.Sqrt(float64(n))))
	for i := int64(2); i <= upper; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func _maxPrime(arr []int64) int64 {
	flag := arr[0]
	for _, i := range arr {
		if i > flag && IsPrime(i) {
			flag = i
		}
	}
	return flag
}

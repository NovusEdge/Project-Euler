package euler

import (
	"fmt"
	"strconv"
	"time"
)

/*
Problem035 answers the problem at : https://projecteuler.net/problem=35

* Problem 35:

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/
func Problem035() {
	start := time.Now()

	primeMap := _makePrimeMap()
	ans := len(primeMap)
	for i := range primeMap {
		k := len(strconv.Itoa(i))
		m := i
		for j := 0; j < k; j++ {
			_, ok := primeMap[m]
			if !ok {
				break
			}
			m = rotate(m)
		}
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 35 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func rotate(n int) int {
	s := strconv.Itoa(n)
	first := rune(s[0])
	chars := []rune(s)
	for i := 0; i < len(s)-1; i++ {
		chars[i] = chars[i+1]
	}
	chars[len(chars)-1] = first

	p, _ := strconv.Atoi(string(chars))
	return p
}

func _makePrimeMap() map[int]bool {
	primes := PrimeSieve(1000000)
	res := make(map[int]bool)
	for _, i := range primes {
		res[i] = true
	}
	return res
}

package euler

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Problem032 answers the problem at : https://projecteuler.net/problem=32

* Problem 32:

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/
func Problem032() {

	start := time.Now()
	ans := 0
	resmap := make(map[int]int)

	for i := 1; i < 2000; i++ {
		for j := 1; j < 2000; j++ {
			temp, _ := strconv.Atoi(fmt.Sprintf("%d%d%d", i, j, i*j))
			if _isPandigital(temp) {
				resmap[i*j] = 1
			}
		}
	}
	for k := range resmap {
		ans += k
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 32 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

func _isPandigital(n int) bool {
	k := "123456789"
	ref := _makeNumStr(n)
	sort.Strings(ref)
	return k == strings.Join(ref, "")
}

func _makeNumStr(n int) []string {
	return strings.Split(strconv.Itoa(n), "")
}

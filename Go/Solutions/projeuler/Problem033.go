package euler

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
Problem033 answers the problem at : https://projecteuler.net/problem=33

* Problem 33:

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

***We shall consider fractions like, 30/50 = 3/5, to be trivial examples.***

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.

*/
func Problem033() {
	res := [][]int{}

	start := time.Now()

	for a := 11; a < 100; a++ {
		for b := 11; b < 100; b++ {
			if check33(a, b) {
				a1, b1, _ := cut(a, b)
				res = append(res, []int{a1, b1})
			}
		}
	}

	end := time.Now()
	fmt.Println("\nAnswer to Problem 33 (Resultant fractions) : ", res)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

func check33(a, b int) bool {
	if a%10 == 0 && b%10 == 0 || a == b {
		return false
	}
	a1, b1, errStr := cut(a, b)
	if errStr != "" {
		return false
	}
	return a1/b1 == a/b
}

func cut(a, b int) (int, int, string) {

	a1 := _makeNumStr(a)
	b1 := _makeNumStr(b)
	for _, i := range a1 {
		if _contains(b1, i) {
			a1[index(a1, i)] = ""
			b1[index(b1, i)] = ""
		}
	}
	res1, _ := strconv.Atoi(strings.Join(a1, ""))
	res2, _ := strconv.Atoi(strings.Join(b1, ""))
	if res2 == 0 {
		return -1, -1, "ZeroDivision"
	}
	return res1, res2, ""
}

func index(arr []string, n string) int {
	for j := 0; j < len(arr); j++ {
		if arr[j] == n {
			return j
		}
	}
	return -1
}

func _contains(s []string, i string) bool {
	for _, j := range s {
		if i == j {
			return true
		}
	}
	return false
}

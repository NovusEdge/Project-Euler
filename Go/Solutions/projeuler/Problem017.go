package euler

import (
	"fmt"
	"time"
)

/*
Problem017 answers the problem at : https://projecteuler.net/problem=17

* Problem 17:


	If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

	If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

	NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.


*/
func Problem017() {
	start := time.Now()
	ans := 0

	for i := 1; i < 1001; i++ {
		ans += len(convTill1000(i))
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 17 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

var numWords = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func convert1to99(n int) (w string) {
	if n < 20 {
		w = numWords[n]
		return
	}
	r := n % 10
	if r == 0 {
		w = numWords[n]
	} else {
		w = numWords[n-r] + numWords[r]
	}
	return
}

func convert100to999(n int) (w string) {
	q := n / 100
	r := n % 100
	w = numWords[q] + "hundred"
	if r == 0 {
		return
	}
	w = w + "and" + convert1to99(r)
	return
}

func convTill1000(n int) (w string) {
	if n < 100 {
		w = convert1to99(n)
		return
	}
	if n == 1000 {
		w = "onethousand"
		return
	}
	w = convert100to999(n)
	return
}

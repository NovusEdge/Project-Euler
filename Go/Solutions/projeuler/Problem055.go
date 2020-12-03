package euler

import (
	"fmt"
	"time"
)

/*
Problem055 answers the problem at : https://projecteuler.net/problem=55
If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome. A n that never forms a palindrome through the reverse and add process is called a Lychrel n. Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a n is Lychrel until proven otherwise. In addition you are given that for every n below ten-thousand, it will either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome. In fact, 10677 is the first n to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?

NOTE: Wording was modified slightly on 24 April 2007 to emphasise the theoretical nature of Lychrel numbers.
*/
func Problem055() {
	start := time.Now()
	ans := 0

	for i := 0; i < 10000; i++ {
		if isLychrel(i) {
			ans++
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 53 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func isLychrel(n int) bool {
	temp := n
	for i := 0; i < 50; i++ {
		temp += reverseNum(temp, 0)
		if IsPall(temp) {
			return false
		}
	}
	return true
}

func reverseNum(n int, res int) int {
	if n != 0 {
		res = res * 10
		res = res + n%10
		n = n / 10
	} else {
		return res
	}
	return reverseNum(n, res)
}

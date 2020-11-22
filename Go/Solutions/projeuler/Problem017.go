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


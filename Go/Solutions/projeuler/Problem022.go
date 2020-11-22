package euler

import (
	"fmt"
	"time"
)

/*
Problem022 answers the problem at : https://projecteuler.net/problem=22

* Problem 22:
		Using names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

		For example, when the list is sorted into alphabetical order, "COLIN", which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, "COLIN" would obtain a score of 938 × 53 = 49714.

		What is the total of all the name scores in the file?
*/
func Problem022() {

	start := time.Now()
	ans := getAns22()
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 21 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

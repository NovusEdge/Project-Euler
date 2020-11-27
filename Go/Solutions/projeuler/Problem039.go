package euler

import (
	"fmt"
	"time"
)

/*
Problem039 answers the problem at : https://projecteuler.net/problem=39

* Problem 39:

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}.
For example, there are exactly three solutions for p = 120:

	{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/
func Problem039() {
	start := time.Now()
	sols := 0
	pFlag := 0

	for p := 1; p <= 1000; p++ {
		temp := getTriSol(p)
		if temp > sols {
			sols = temp
			pFlag = p
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 39 : %d\n", pFlag)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


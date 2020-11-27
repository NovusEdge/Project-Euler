package euler

import (
	"fmt"
	"time"
)

/*
Problem020 answers the problem at : https://projecteuler.net/problem=20

* Problem 20:
		Find the sum of the digits in the number 100!
*/
func Problem020() {
	start := time.Now()
	ans := DigitalRoot(Factorial(100))
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 20 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

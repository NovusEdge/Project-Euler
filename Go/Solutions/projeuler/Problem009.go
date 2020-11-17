package euler

import (
	"fmt"
	"math"
	"time"
)

/*
Problem009 answers the problem at : https://projecteuler.net/problem=9

* Problem 9:
		There exists exactly one Pythagorean triplet for which a + b + c = 1000 (a < b < c).
		Find the product abc.
*/
func Problem009() {

	start := time.Now()
	ans := get9()
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 9 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func isTriplet(a, b, c int) bool {
	return c*c == a*a+b*b && a < b && b < c
}

func get9() int {
	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			k := math.Sqrt(float64(i*i + j*j))

			if isTriplet(i, j, int(k)) && i+j+int(k) == 1000 {
				return i * j * int(k)
			}
		}
	}
	return 0
}

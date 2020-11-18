package euler

import (
	"fmt"
	"strconv"
	"time"
)

/*
Problem038 answers the problem at : https://projecteuler.net/problem=38

* Problem 38:

Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

***What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?***
*/
func Problem038() {
	start := time.Now()
	maxProd := 1

	for i := 1; i < 10000; i++ {
		temp := getProduct(i)
		if maxProd < temp {
			maxProd = temp
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 38 : %d\n", maxProd)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func getProduct(n int) int {
	i := 1
	for true {
		temp := ""
		for j := 1; j <= i; j++ {
			k := strconv.Itoa(n * j)
			temp += k
		}
		if len(temp) == 9 {
			if isPallStr(temp) {
				p, _ := strconv.Atoi(temp)
				return p
			} else {
				return 0
			}
		}
		if len(temp) > 9 {
			return 0
		}
		i++
	}
	return 0
}

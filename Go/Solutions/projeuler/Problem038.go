package euler

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
	ans := 2

	for i := 2; i < 10000; i++ {
		temp := panProd(i)
		if ans < temp {
			ans = temp
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 38 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func panProd(n int) int {
	i := 0
	for true {
		temp := process(n, i)
		if len(temp) == 9 {
			if _pallStr38(temp) {
				res, _ := strconv.Atoi(temp)
				return res
			}
			return 0
		}
		if len(temp) > 9 {
			return 0
		}
		i++
	}
	return 0
}

func process(n, upper int) string {
	temp := ""

	for i := 1; i <= upper; i++ {
		temp += strconv.Itoa(n * i)
	}
	return temp
}
func _pallStr38(s string) bool {
	k := []string{}
	for _, i := range s {
		k = append(k, string(i))
	}
	sort.Strings(k)
	return strings.Join(k, "") == "123456789"
}

package euler

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

/*
Problem016 answers the problem at : https://projecteuler.net/problem=16

* Problem 16:

		2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
		What is the sum of the digits of the number 2^1000?

*/
func Problem016() {
	start := time.Now()
	num := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil)
	ans := DigitalRoot(num)
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 16 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

/*
DigitalRoot : Returns the sum of all digits of a number
*/
func DigitalRoot(n *big.Int) int64 {
	numStr := n.String()
	res := 0
	for _, i := range numStr {
		j, _ := strconv.Atoi(string(i))
		res += j
	}
	return int64(res)
}

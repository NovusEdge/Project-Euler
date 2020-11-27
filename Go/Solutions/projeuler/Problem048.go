package euler

import (
	"fmt"
	"math/big"
	"time"
)

/*
Problem048 answers the problem at : https://projecteuler.net/problem=48

* Problem 48:

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/
func Problem048() {

	start := time.Now()
	ans := big.NewInt(0)

	for i := 1; i < 1001; i++ {
		k := big.NewInt(int64(i))
		k.Exp(k, k, nil)
		ans.Add(ans, k)
	}
	ans.Mod(ans, big.NewInt(10000000000))

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 42 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

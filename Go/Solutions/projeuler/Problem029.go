package euler

import (
	"fmt"
	"math/big"
	"time"
)

/*
Problem029 answers the problem at : https://projecteuler.net/problem=29

* Problem 29:


Consider all integer combinations of *a^b* for 2 ≤ *a* ≤ 5 and 2 ≤ *b* ≤ 5:

    2^2=4, 2^3=8, 2^4=16, 2^5=32
    3^2=9, 3^3=27, 3^4=81, 3^5=243
    4^2=16, 4^3=64, 4^4=256, 4^5=1024
    5^2=25, 5^3=125, 5^4=625, 5^5=3125

If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ *a* ≤ 100 and 2 ≤ *b* ≤ 100?
*/
func Problem029() {

	makeSet := func(arr []*big.Int) []*big.Int {
		temp := []*big.Int{}
		for _, i := range arr {
			if !inArr(temp, i) {
				temp = append(temp, i)
			}
		}
		return temp
	}
	nums := []*big.Int{}

	start := time.Now()
	for i := 2; i < 101; i++ {
		for j := 2; j < 101; j++ {
			nums = append(nums, pow(int64(i), int64(j)))
		}
	}
	ans := len(makeSet(nums))
	end := time.Now()

	fmt.Println("\nAnswer to Problem 29 : ", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


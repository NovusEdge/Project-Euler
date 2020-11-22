package euler

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

/*
Problem042 answers the problem at : https://projecteuler.net/problem=42

* Problem 42:


The nth term of the sequence of triangle numbers is given by, t(n) = n(n+1)/2; so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using **words.txt** (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/
func Problem042() {
	start, ans := time.Now(), 0

	TNums := genTriSeq()
	fileObj, _ := ioutil.ReadFile("./p042_words.txt")
	data := strings.Split(string(fileObj), `,`)
	trimQuotes(&data)
	k := mapToScore(data)
	for _, i := range k {
		p, _ := binarySearch(TNums, i)
		if p != -1 {
			ans++
		}
	}

	end := time.Now()
	fmt.Printf("\nAnswer to Problem 42 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

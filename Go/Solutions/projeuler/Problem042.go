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
	fmt.Printf("\nAnswer to Problem 41 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}

func trimQuotes(data *[]string) {
	for i := 0; i < len(*data); i++ {
		p := (*data)[i]
		(*data)[i] = p[1 : len(p)-1]
	}
}

func mapToScore(data []string) (res []int64) {
	for _, i := range data {
		res = append(res, int64(numScore(i)))
	}
	return
}

func genTriSeq() (res []int64) {
	for i := 1; i < 1000; i++ {
		res = append(res, TriangleNum(int64(i)))
	}
	return
}

func binarySearch(a []int64, search int64) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}
	searchCount++
	return
}

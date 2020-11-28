package euler

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

//returns the nth fibonacci term if it's even else returns 0 (used in Problem 2)
func _evenfibb(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a += b
		a, b = b, a
	}
	if a%2 == 0 {
		return a
	}
	return 0
}

//Factors returns an array of the factors of a number
func Factors(n int64) (res []int64) {
	upper := int64((math.Sqrt(math.Abs(float64(n)))))
	for i := int64(1); i < upper; i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, int64(n/i))
		}
	}
	return
}

//IsPrime reports if a number is a prime number or not
func IsPrime(n int64) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}
	upper := int64(math.Trunc(math.Sqrt(float64(n))))
	for i := int64(2); i <= upper; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//reports the largest prime number that is in the array [arr] (used in problem 3)
func _maxPrime(arr []int64) int64 {
	flag := arr[0]
	for _, i := range arr {
		if i > flag && IsPrime(i) {
			flag = i
		}
	}
	return flag
}

//IsPall checks if a number is a pallindrome
func IsPall(n int) bool {
	s := fmt.Sprintf("%d", int(math.Abs(float64(n))))
	return s == Reverse(s)
}

//Reverse returns a reversed string of [s]
func Reverse(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

//Max reports the max-element in [arr]
func Max(arr []int) int {
	flagInt := arr[0]
	for _, i := range arr {
		if i > flagInt {
			flagInt = i
		}
	}
	return flagInt
}

//GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

//LCM : find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// reports the sum of the first *n* squares (used in problem 6)
func squareSum(n int) int {
	k := 0
	for i := 1; i <= n; i++ {
		k += i * i
	}
	return k
}

//reports the square of the sum of first n numbers (used in problem 6)
func sumSquare(n int) int {
	var k int
	k = n * (n + 1) / 2
	return k * k
}

//PrimeSieve implements the sieve of Eratosthenes
func PrimeSieve(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] == true {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}
	return primes
}

//checks if a triplet of integers forms a right triangle
func isTriplet(a, b, c int) bool {
	return c*c == a*a+b*b && a < b && b < c
}

func getAns9() int {
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

//Sum reports the sum of all elements of [arr]
func Sum(arr []int) (s int) {
	for _, i := range arr {
		s += i
	}
	return
}

func highTriangle(n int64) bool {
	return len(Factors(n)) > 500
}

//TriangleNum reports the nth triangle number
func TriangleNum(n int64) int64 {
	return n * (n + 1) / 2
}

func getAns12() int64 {
	var i int64 = 1
	for !highTriangle(TriangleNum(i)) {
		i++
	}
	return TriangleNum(i)
}

func getAns13(arr []string) string {
	res := big.NewInt(0)
	var temp big.Int
	for i := 0; i < len(arr); i++ {
		intObj, _ := temp.SetString(arr[i], 10)
		res.Add(res, intObj)
	}
	return res.String()[:10]
}

//generates collatz chain for a number [n] terminates at 1
func collatz(n int) []int {
	resArr := []int{n}
	i := n
	for i > 1 {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = 3*i + 1
		}
		resArr = append(resArr, i)
	}
	return resArr
}

func getAllSeq() (resArr []int) {
	for i := 2; i < 1000000; i++ {
		resArr = append(resArr, len(collatz(i)))
	}
	return
}

func getMaxFromSeq(arr []int) int {
	flagLen := len(collatz(2))
	flagElem := 2
	for i := 0; i < len(arr); i++ {
		if arr[i] > flagLen {
			flagElem = i + 2
			flagLen = arr[i]
		}
	}
	return flagElem
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

var numWords = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// converts numbers from 1-99 to their "words"
func convert1to99(n int) (w string) {
	if n < 20 {
		w = numWords[n]
		return
	}
	r := n % 10
	if r == 0 {
		w = numWords[n]
	} else {
		w = numWords[n-r] + numWords[r]
	}
	return
}

// converts numbers from 100-999 to their "words"
func convert100to999(n int) (w string) {
	q := n / 100
	r := n % 100
	w = numWords[q] + "hundred"
	if r == 0 {
		return
	}
	w = w + "and" + convert1to99(r)
	return
}

// converts all the numbers from 1-1000 to their "words"
func convTill1000(n int) (w string) {
	if n < 100 {
		w = convert1to99(n)
		return
	}
	if n == 1000 {
		w = "onethousand"
		return
	}
	w = convert100to999(n)
	return
}

//Factorial reports the factorial of a number
func Factorial(n int64) *big.Int {
	return big.NewInt(1).MulRange(2, n)
}

//reports the score of a name/word
func numScore(name string) (res int) {
	for _, i := range []rune(name) {
		res += int(i) - 64
	}
	return
}

func getAns22() (res int) {
	nameFile, _ := ioutil.ReadFile("p022_names.txt")
	names := strings.Split(string(nameFile[:]), ",")
	sort.Strings(names)
	for i := 0; i < len(names); i++ {
		res += numScore(names[i][1:len(names[i])-1]) * (i + 1)
	}
	return
}

//checks if a number is an abundant number or not
func isAbundant(n int) bool {
	var res int64
	k := Factors(int64(n))
	for _, i := range k {
		res += i
	}
	return res-int64(n) > int64(n)
}

func getAbundant() (res []int) {
	for i := 1; i < 28124; i++ {
		if isAbundant(i) {
			res = append(res, i)
		}
	}
	return
}

func getPerms(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1]
	remStr := str[1:]
	perms := getPerms(remStr)
	allPerms := make([]string, 0)

	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}
	return allPerms
}

func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]

	return start + char + end
}

/*
Fibonacci :Reports the nth term in the fibonacci series

* Fibonacci(0) --> 1

* Fibonacci(1) --> 1

* Fibonacci(2) --> 2

* ...
*/
func Fibonacci(n int64) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	limit := big.NewInt(n)
	i := big.NewInt(0)

	for i.Cmp(limit) < 0 {
		a.Add(a, b)
		a, b = b, a
		i.Add(i, big.NewInt(1))
	}
	return a
}

func numLen(n *big.Int) int {
	k := n.String()
	return len(k)
}

func formQuad(a, b int) (int, string) {
	counter, i := 0, 0
	for true {
		exprn := i*i + a*i + b
		if IsPrime(int64(exprn)) {
			counter++
			i++
		} else {
			return counter, fmt.Sprintf("n**2 + %dn + %d", a, b)
		}
	}
	return 0, ""
}

func dictMax(m map[int]string) int {
	flag := 0
	for k := range m {
		if k > flag {
			flag = k
		}
	}
	return flag
}

func pow(i, j int64) (res *big.Int) {
	a := big.NewInt(i)
	b := big.NewInt(j)
	res = a.Exp(a, b, nil)
	return
}

func inArr(arr []*big.Int, e *big.Int) bool {
	for _, i := range arr {
		if i.Cmp(e) == 0 {
			return true
		}
	}
	return false
}

//Digits returns an array of digits of [n]
// func Digits(n int) (res []int) {
// 	numStr := strconv.Itoa(n)
// 	for _, i := range numStr {
// 		k, _ := strconv.Atoi(string(i))
// 		res = append(res, k)
// 	}
// 	return
// }

//Digits returns an array of digits of [n]
func Digits(n int) (res []int) {
	temp := n
	for temp > 0 {
		res = append(res, temp%10)
		temp /= 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return
}

func fifthPowSum(digits []int) (res int) {
	for _, i := range digits {
		res += i * i * i * i * i
	}
	return
}

func _isPandigital(n string) bool {
	k := "123456789"
	ref := strings.Split(n, "")
	sort.Strings(ref)
	return k == strings.Join(ref, "")
}

func _makeNumStr(n int) []string {
	return strings.Split(strconv.Itoa(n), "")
}

func check33(a, b int) bool {
	if a%10 == 0 && b%10 == 0 || a == b {
		return false
	}
	a1, b1, errStr := cut(a, b)
	if errStr != "" {
		return false
	}
	return a1/b1 == a/b
}

func cut(a, b int) (int, int, string) {

	a1 := _makeNumStr(a)
	b1 := _makeNumStr(b)
	for _, i := range a1 {
		if _contains(b1, i) {
			a1[index(a1, i)] = ""
			b1[index(b1, i)] = ""
		}
	}
	res1, _ := strconv.Atoi(strings.Join(a1, ""))
	res2, _ := strconv.Atoi(strings.Join(b1, ""))
	if res2 == 0 {
		return -1, -1, "ZeroDivision"
	}
	return res1, res2, ""
}

func index(arr []string, n string) int {
	for j := 0; j < len(arr); j++ {
		if arr[j] == n {
			return j
		}
	}
	return -1
}

func _contains(s []string, i string) bool {
	for _, j := range s {
		if i == j {
			return true
		}
	}
	return false
}

func _fact34(n int) (res int) {
	res = 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return
}

func check34(n int) bool {
	temp := 0
	for _, i := range Digits(n) {
		temp += _fact34((i))
	}
	return temp == n
}

func rotate(n int) int {
	digits := Digits(n)
	res, l := make([]int, len(digits)), len(digits)
	copy(res, digits)

	for i := 0; i < l-1; i++ {
		res[i] = res[i+1]
	}

	res[l-1] = digits[0]
	return formNum(res)
}

func _makePrimeMap() map[int]bool {
	primes := PrimeSieve(1000000)
	res := make(map[int]bool)
	for _, i := range primes {
		res[i] = true
	}
	return res
}

func check35(n int) bool {
	ref, temp := n, rotate(n)

	for temp != ref {
		_, ok := primeMap[temp]
		if !ok {
			return false
		}
		temp = rotate(temp)
	}
	return true
}

func isPallStr(s string) bool {
	return s == _reverse(s)
}

func _reverse(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

func bin(n int64) string {
	return strconv.FormatInt(n, 2)
}

func isTruncL(n int64) bool {
	s := []byte(strconv.FormatInt(n, 10))

	for i := 0; i < len(s); i++ {
		k, _ := strconv.Atoi(string(s[i:]))
		if !IsPrime(int64(k)) {
			return false
		}
	}
	return true
}

func isTruncR(n int64) bool {
	s := []byte(strconv.FormatInt(n, 10))

	for i := len(s); i > 0; i-- {
		k, _ := strconv.Atoi(string(s[:i]))
		if !IsPrime(int64(k)) {
			return false
		}
	}
	return true
}

func panProd(n int) int {
	i := 0
	for true {
		temp := process38(n, i)
		if len(temp) == 9 && _pallStr38(temp) {
			res, _ := strconv.Atoi(temp)
			return res
		}
		if len(temp) > 9 {
			return 0
		}
		i++
	}
	return 0
}

func process38(n, upper int) string {
	temp := ""

	for i := 1; i <= upper; i++ {
		temp += strconv.Itoa(n * i)
	}
	return temp
}

//_pallStr38 checks for 1-9 pallindrome
func _pallStr38(s string) bool {
	k := []string{}
	for _, i := range s {
		k = append(k, string(i))
	}
	sort.Strings(k)
	return strings.Join(k, "") == "123456789"
}

func getTriSol(p int) int {
	soln := 0

	for i := 1; i < 400; i++ {
		for j := 1; j < 400; j++ {
			k := p - i - j
			if _isIntRT(i, j, k) && k > 0 {
				soln++
			}
		}
	}
	return soln
}

// checks if a triangle with sides, a, b, c (c as the hypotenuse) is a right triangle (integer right triangle)
func _isIntRT(a, b, c int) bool {
	return c*c == a*a+b*b && a < b && b < c
}

func getPrimePerm(n int) (res int) {
	perms := permutations(Digits(n))
	for _, i := range perms {
		num := formNum(i)
		if IsPrime(int64(num)) && res < num {
			res = num
		}
	}
	return
}

func formNum(digits []int) (res int) {
	l := len(digits)
	res += digits[0] * int(math.Pow(10, float64(l)))

	for i := 1; i < l; i++ {
		p := int(math.Pow(10, float64(l-i)))
		res += digits[i] * p
	}
	res = res / 10
	return
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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

/* Reports the nth **Triangle** number */
func tNum(n int64) int64 {
	return n * (n + 1) / 2
}

/* Reports the nth **Pentagonal** number */
func pNum(n int64) int64 {
	return n * (3*n - 1) / 2
}

/* Reports the nth **Hexagonal** number */
func hNum(n int64) int64 {
	return n * (2*n - 1)
}

/* Checks if 2 numbers have the same digits (used in problem 52) */
func sameDigits(a int, b int) bool {
	a1 := Digits(a)
	b1 := Digits(b)

	if len(a1) != len(b1) {
		return false
	}

	sort.Ints(a1[:])
	sort.Ints(b1[:])

	for i := 0; i < len(a1); i++ {
		if a1[i] != b1[i] {
			return false
		}

	}
	return true
}

//C reports the number of combinations i.e. nCr
func C(n, r int) *big.Int {
	t1, t2, t3 := Factorial(int64(n)), Factorial(int64(r)), Factorial(int64(n-r))
	t2.Mul(t2, t3)
	res := t1.Div(t1, t2)
	return res
}

# from itertools import islice,count
# from math import floor,sqrt
from functools import reduce
from time import time
from math import ceil

''' 
Problem 69: https://projecteuler.net/problem=69

Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number of numbers less than n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.
n 	Relatively Prime 	φ(n) 	 n/φ(n)
2 	1 	                1 	     2
3 	1,2 	            2        1.5
4 	1,3 	            2     	 2
5 	1,2,3,4 	        4        1.25
6 	1,5 	            2 	     3
7 	1,2,3,4,5,6 	    6 	     1.1666...
8 	1,3,5,7 	        4 	     2
9 	1,2,4,5,7,8 	    6 	     1.5
10 	1,3,7,9 	        4 	     2.5

It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.

Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
'''



''' 
# In number theory, Euler's totient function counts the positive integers up to a given integer n that are relatively prime to n.

# It is written using the Greek letter phi as φ(n) or ϕ(n), and may also be called Euler's phi function.

In other words, it is the number of integers k in the range 1 ≤ k ≤ n for which the greatest common divisor gcd(n, k) is equal to 1.[2][3] The integers k of this form are sometimes referred to as totatives of n. 
'''

def factors(n):    
    return set(reduce(list.__add__, 
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))
    
def primeSieve(n):
    sieve = [True] * n
    for i in range(3, int(n**0.5)+1, 2):
        if sieve[i]:
            sieve[i*i:: 2*i] = [False] * ((n-i*i - 1) // (2 * i) + 1)
    return [2] + [i for i in range(3, n, 2) if sieve[i]]

pSieve = primeSieve(1_000_001)

def phi(n):
    prodt = reduce(float.__mul__, [1-(1/pSieve[i]) for i in range(n)])
    return ceil(n*prodt)
# https://en.wikipedia.org/wiki/Euler%27s_totient_function

start = time()
for i in range(1, 10):
    print(phi(i))
# needs work ...
ans = phi(6)

if __name__ == "__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

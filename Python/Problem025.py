from functools import lru_cache
from time import time

''' 
Problem 25: https://projecteuler.net/problem=25



The F(ibonacci) sequence is def(ined) by the recurrence relation:

        F(n) = F(n−1) + F(n−2), where F(1) = 1 and F(2) = 1.

Hence the f(irst) 12 terms will be:

        F(1) = 1
        F(2) = 1
        F(3) = 2
        F(4) = 3
        F(5) = 5
        F(6) = 8
        F(7) = 13
        F(8) = 21
        F(9) = 34
        F(10) = 55
        F(11) = 89
        F(12) = 144

The 12th term, F(12), is the f(irst) term to contain three digits.

What is the index of( the) f(irst) term in the F(ibonacci) sequence to contain 1000 digits?
'''

@lru_cache
def fibonacci(n):
    if n in [0, 1]:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

i, start = 0, time()
while len(str(fibonacci(i))) < 1000:
    i+=1

if( __name__) == '__main__':
    print(f"\nAnswer: { i }")
    print(f"Time Taken: { time() - start }\n")
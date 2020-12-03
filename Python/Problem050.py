from time import time
from math import sqrt; from itertools import count, islice

''' 
Problem 50: https://projecteuler.net/problem=50

The prime 41, can be written as the sum of six consecutive primes:
                    41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
'''

def prime(n):
    return n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

def primeSieve(n):
    sieve = [True] * n
    for i in range(3,int(n**0.5)+1,2):
        if sieve[i]:
            sieve[i*i::2*i]=[False]*((n-i*i-1)//(2*i)+1)
    return [2] + [i for i in range(3,n,2) if sieve[i]]


start = time() 
pSieve = primeSieve(1000000)
ans = 0

for i in range(1, len(pSieve)):
    k = sum(pSieve[:i+1])
    if k >= 1000000: break
    if prime(k): ans = k

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
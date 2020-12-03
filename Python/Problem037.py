from time import time
from math import sqrt; from itertools import count, islice

''' 
Problem 37: https://projecteuler.net/problem=37

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
'''

prime = lambda n: n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

def isTrunc_L(n):
    i = 0
    while i < len(str(n)):
        if not prime(int(str(n)[i:])):
            return False
        i += 1
    return True

def isTrunc_R(n):
    i = len(str(n))
    while i > 0:
        if not prime(int(str(n)[:i])):
            return False
        i -= 1
    return True

start = time()
i, counter = 10, 0
ans = 0

while counter < 11:
    if isTrunc_R(i) and isTrunc_L(i):
        ans += i
        counter += 1
        print(i)
    i += 1

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")


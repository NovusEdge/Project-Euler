from itertools import islice,count
from math import floor,sqrt
from functools import reduce
from time import time

''' 
Problem 3: https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
'''

def prime(n):
    return n > 1 and all(n % i for i in islice(count(2), int(sqrt(n) - 1)))

def factors(n):    
    return set(reduce(list.__add__, 
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))
    
start = time()
ans = max(filter(lambda x: prime(x), factors(600851475143)))


if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
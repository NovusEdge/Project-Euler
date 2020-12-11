from time import time
from math import lcm
from functools import reduce

''' 
Problem 5: https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
'''

start = time()
ans = reduce(lcm, range(1, 21))

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

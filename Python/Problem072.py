from time import time
from math import gcd
from functools import reduce

''' 
Problem 72: https://projecteuler.net/problem=72

Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.

If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:

    1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

It can be seen that there are 21 elements in this set.

How many elements would be contained in the set of reduced proper fractions for d ≤ 1,000,000?
'''

def factors(n):    
    return set(reduce(list.__add__, 
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))

# isReduced = lambda a, b: a < b and gcd(a, b) == 1
# n_set, d_set = set(), set()

def main():
    
    
    pass

'''1/8, 3/8, 1/7, 2/7, 3/7, 1/6, 1/5, 2/5, 1/4, 1/3, 1/2, 2/3, 3/4, 3/5, 4/5, 5/6, 4/7, 5/7, 6/7, 5/8, 7/8'''

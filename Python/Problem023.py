from math import sqrt
from time import time
from functools import reduce

''' 
Problem 23: https://projecteuler.net/problem=23

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
'''

def factors(n):    
    return set(reduce(list.__add__, 
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))

start = time()
temp_list = list(filter(lambda x: sum(factors(x))-x > x, range(1, 28124)))
boolMap = dict()

for i in range(1, 28124): boolMap[i] = True

for i in temp_list:
    for j in temp_list:
        boolMap[i+j] = False

ans = sum([i for i in boolMap if boolMap[i]])

if __name__ == "__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time taken:{ time() - start }\n")

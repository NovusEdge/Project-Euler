from math import floor, sqrt
from time import time

''' 
Problem 9: https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
                    
                    a**2 + b**2 = c**2

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.
'''

def isTriplet(a, b, c): return a < b < c and a**2 + b**2 == c**2 and a+b+c == 1000


def main():
    for i in range(1000):
        for j in range(1000):
            k = (i**2 + j**2) ** 0.5
            if isTriplet(i, j, int(k)):
                return i*j*int(k)


start = time()
ans = main()

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

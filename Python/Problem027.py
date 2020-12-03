from itertools import islice, count
from math import floor, sqrt
from time import time

''' 
Problem 27: https://projecteuler.net/problem=27

Considering quadratics of the form:
		n^2 + a*n + b , where |a| < 1000 and |b| <= 1000

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
'''


def prime(n):
    if n == 0 or n == 1:
        return None
    if n == 2:
        return True
    return all([not(n % i == 0) for i in islice(count(2), floor(sqrt(n)))])


def checkQuad(a, b):
    i = 0
    counter = 0
    while True:
        exprn = pow(i, 2) + a*i + b
        if prime(exprn):
            counter += 1
            i += 1
        else:
            return (counter, "n**2 + {}n + {}".format(a, b))


def main():
    maindict = {}
    for a in range(-1000, 1001):
        for b in range(-1000, 1000):
            try:
                k = checkQuad(a, b)
                maindict[k[0]] = k[1]
            except:
                pass
    return maindict[max(maindict)]


if __name__ == "__main__":
    start = time()
    print(f"\nThe quadratic expression is:{ main() }")
    print(f"Time Taken:{ time() - start }\n")

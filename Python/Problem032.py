from time import time

''' 
Problem 32: https://projecteuler.net/problem=32

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
'''

isPan = lambda n: sorted(n) == ['1', '2', '3', '4', '5', '6', '7', '8', '9']


def main():
    resArr = set()
    for i in range(1, 2000):
        for j in range(1, 2000):
            if isPan(f"{ i }{ j }{ i*j }"):
                resArr.add(i * j)
    return sum(resArr)


if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
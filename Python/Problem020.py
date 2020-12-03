from math import factorial
from time import time

''' 
Problem 20: https://projecteuler.net/problem=20

Find the sum of the digits in the number 100!
'''

digits = lambda n: map(int, list(str(n)))

start = time()
ans = sum(digits(factorial(100)))

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

from time import time

''' 
Problem 6: https://projecteuler.net/problem=6

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
'''

start = time()
sqSum = lambda n: int(n * (n+1) * (2*n + 1) / 6)
sumSq = lambda n: int(n * (n+1) / 2) ** 2

ans = sumSq(100) - sqSum(100)

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
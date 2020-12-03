from time import time
from math import factorial

''' 
Problem 34: https://projecteuler.net/problem=34

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
'''

digitFact = lambda n: sum(map(lambda x: factorial(int(x)), str(n)))

ans = 0
for i in range(3, 100000):
    if digitFact(i) == i:
        ans += i

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

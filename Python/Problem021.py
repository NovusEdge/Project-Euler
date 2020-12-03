from functools import reduce
from time import time

''' 
Problem 21: https://projecteuler.net/problem=21

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
'''

def d(n):
    return sum(set(reduce(list.__add__, 
                ([i, n // i] for i in range(1, int(n**0.5) + 1) if n%i == 0))) - {n})
                
def main():
    check = lambda a, b : d(a) == b and d(b) == a and a != b
    ans = 0
    
    for i in range(1, 10000):
        for j in range(1, i+1):
            if check(i, j):
                ans += i+j
            
    return ans

if __name__ == "__main__":
  start = time()
  print(f"\nAnswer : { main() }")
  print(f"Time Taken : { time() - start }\n")

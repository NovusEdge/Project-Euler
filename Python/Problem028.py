from time import time

''' 
Problem 28: https://projecteuler.net/problem=28

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

                    21 22 23 24 25
                    20  7  8  9 10
                    19  6  1  2 11
                    18  5  4  3 12
                    17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
'''

def getSpiralSum( n ):
    if n == 0: return 1
    if n == 1: return 25
    if n == 2: return 101
    
    return 4*(2*n + 1)**2 - (12*n) + getSpiralSum(n - 1)

start = time()
ans = getSpiralSum(500)

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
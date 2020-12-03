from time import time
from math import factorial as fact
from itertools import combinations_with_replacement as cwr

''' 
Problem 92: https://projecteuler.net/problem=92

A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.

For example,

                    44 → 32 → 13 → 10 → 1 → 1
                    85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?
'''

digitSquare = lambda n: sum([i**2 for i in map(int, str(n))])

def checkChain( n ):
    temp = n
    while True:
        temp = digitSquare( temp )
        
        if temp == 89: 
            return True
        if temp == 1 or temp == 0: 
            return False

start = time()
numRef = '1234567890'
counter = 0
combLis = []
for i in range(1, 10):
    if i < 8:
        combs = cwr(numRef, i)
    else: combs = cwr(numRef, 8)
    counter = int(fact(len(numRef[:i])+i-1) / fact(i) / fact(len(numRef[:i])-1)) * len(numRef[:i])
    
# ans = sum(map(lambda x: fact(len(str(x))), filter(checkChain, combLis)))
# # (n+r-1)! / r! / (n-1)! needs work ...

if __name__ == '__main__':
    print(f"\nAnswer: { counter }")
    print(f"Time Taken: { time() - start }\n")

# 8581146

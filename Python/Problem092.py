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

checked = dict()
sumDigitSq = lambda n: sum([i**2 for i in map(int, list(str(n)))])

def chainEnd(n):
    try:
        if checked[''.join(sorted(str(ans)))]: return 89
    except:
        temp = n
        while True:
            temp = sumDigitSq(temp)
            if temp == 89 or temp == 1:
                checked[''.join(sorted(str(ans)))] = (temp == 89)
                break
            
ans = 0
start = time()

for i in range(10*10**6):
    if chainEnd(ans) == 89:
        if checked[''.join(sorted(str(ans)))]:
            continue
        ans += 1
        
if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
# 8581146

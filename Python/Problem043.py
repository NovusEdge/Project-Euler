from time import time
from itertools import permutations as perm
''' 
Problem 43: https://projecteuler.net/problem=43


The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

    d2 d3 d4 = 406 is divisible by 2
    d3 d4 d5 = 063 is divisible by 3
    d4 d5 d6 = 635 is divisible by 5
    d5 d6 d7 = 357 is divisible by 7
    d6 d7 d8 = 572 is divisible by 11
    d7 d8 d9 = 728 is divisible by 13
    d8 d9 d10 = 289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
'''

def getSubStrings(n):
    subStrs = []
    for i in range(1, 8):
        subStrs.append(str(n)[i:i+3])
        
    return list(map(int, subStrs))

def check(n):
    subStrs = zip(getSubStrings(n), [2, 3, 5, 7, 11, 13, 17])
        
    for i, j in subStrs:
        if i%j != 0:
            return False
        
    return True

start = time() 
ans = 0
panNums = map(lambda x: int(''.join(x)), perm('1234567890', 10))

for i in panNums:
    if check(i):
        ans += i

if __name__=="__main__":    
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
from itertools import permutations as perm
from time import time

'''
Problem 206:  https://projecteuler.net/problem=206

Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,

where each “_” is a single digit.
'''

isSquare = lambda x: int(x**0.5) == x ** 0.5

def makeNumber( nums ):
    return int( "1{}2{}3{}4{}5{}6{}7{}8{}9{}0".format(*nums) )

def main():
    perms = list(perm('1234567890', 9))[::-1]

    for i in perms:
        if isSquare(makeNumber(i)):
            return int(makeNumber(i) ** 0.5)
# needs work
if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")

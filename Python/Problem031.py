from itertools import combinations as comb
from time import time
from functools import reduce

''' 
Problem 31: https://projecteuler.net/problem=31

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
'''

"""
*** 
!!! THIS SOLUTION NEEDS WORK !!! 
***
"""


ans = 0
start = time()
coins = {
    "1p" : 1,
    "2p" : 2,
    "3p" : 3,
    "5p" : 5,
    "10p": 10,
    "20p": 20,
    "50p": 50,
    "1e" : 100,
    "2e" : 200
}
mainComb = [list(comb(coins.keys(), i)) for i in range(1, 10)]




if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")


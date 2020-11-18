from itertools import combinations as comb
from time import time
from functools import reduce

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


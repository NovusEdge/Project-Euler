from time import time
from math import gcd
from functools import reduce

lcm = lambda a, b: a*b//gcd(a, b)

start = time()
ans = reduce(lcm, range(1, 21))

if __name__ == '__main__':
    print(f"Answer: { ans }")
    print(f"Time Taken: { time() - start }")
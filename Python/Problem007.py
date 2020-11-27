from time import time
from itertools import islice,count
from math import floor,sqrt

def prime(n):
    return n > 1 and all(n % i for i in islice(count(2), int(sqrt(n) - 1)))

start = time()
pLis, i = [], 2

while len(pLis) < 10001:
    if prime(i): pLis.append(i)
    i += 1
    
ans = pLis[-1]

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
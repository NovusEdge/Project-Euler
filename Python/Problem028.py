from time import time
from functools import lru_cache

@lru_cache
def getSpiralSum( n ):
    if n == 1 or n == 0: return n
    if n == 3: return 25
    if n == 5: return 61
    
    return 4*(2*n + 1)**2 - (12*n) + getSpiralSum(n - 1)

start = time()
ans = getSpiralSum(500)

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
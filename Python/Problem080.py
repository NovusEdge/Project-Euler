from time import time
from decimal import *
from itertools import islice, count

getcontext().prec = 100
print(len(str(Decimal(2).sqrt() - int(2**0.5))))

def rootDigitalSum(n):
    temp = Decimal(n).sqrt() - int(n**0.5)
    return sum(map(int, list(str(temp)[2:])))

perfSquares = [i**2 for i in range(11)]

ans = 0
start = time()

for i in range(0, 101):
    if i not in perfSquares:
        print(rootDigitalSum(i), i)
        ans += rootDigitalSum(i)
    
#  needs work ...

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")


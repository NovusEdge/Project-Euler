from time import time
from math import factorial
from itertools import combinations_with_replacement as cwr

digitSquare = lambda n: sum([i**2 for i in map(int, str(n))])

def getSquareChain( n ):
    temp, chain = n, [ n ]
    while True:
        temp = digitSquare( temp )
        chain.append( temp )
        if temp == 1 or temp == 89 or temp == 0:
            chain.append(temp)
            return chain

start = time()
numRef = '1234567890'
combs = list(cwr(numRef, 8)) # needs work ...
combs = map(lambda x: int(''.join(sorted(x))), combs)
combs = list(filter(lambda x: getSquareChain(x)[-1]==89, combs))
ans = len(combs)

if __name__ == '__main__':
    print(f"Answer: { ans }")
    print(f"Time Taken: { time() - start }")
# 8581146
# 538063582
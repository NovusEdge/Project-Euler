from time import time
from math import sqrt; from itertools import count, islice

prime = lambda n: n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

def isTrunc_L(n):
    i = 0
    while i < len(str(n)):
        if not prime(int(str(n)[i:])):
            return False
        i += 1
    return True

def isTrunc_R(n):
    i = len(str(n))
    while i > 0:
        if not prime(int(str(n)[:i])):
            return False
        i -= 1
    return True

start = time()
i, counter = 10, 0
ans = 0

while counter < 11:
    if isTrunc_R(i) and isTrunc_L(i):
        ans += i
        counter += 1
    i += 1

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")


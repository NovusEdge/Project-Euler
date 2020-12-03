from time import time
from math import sqrt; from itertools import count, islice

def prime(n):
    return n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

def primeSieve(n):
    sieve = [True] * n
    for i in range(3,int(n**0.5)+1,2):
        if sieve[i]:
            sieve[i*i::2*i]=[False]*((n-i*i-1)//(2*i)+1)
    return [2] + [i for i in range(3,n,2) if sieve[i]]


start = time() 
ans = 0
pSieve = primeSieve(1000000)


if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

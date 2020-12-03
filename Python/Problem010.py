from math import sqrt; from itertools import count, islice
from time import time

def primeSieve(n):
    sieve = [True] * n
    for i in range(3, int(n**0.5) + 1, 2):
        if sieve[i]:
            sieve[i*i : : 2*i] = [False] * ((n - i*i - 1) // (2 * i) + 1)
    return [2] + [i for i in range(3, n, 2) if sieve[i]]

start = time()
ans = sum(primeSieve(2 * 10**6))

if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
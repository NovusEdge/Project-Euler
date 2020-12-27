from time import time


'''
Problem 46: https://projecteuler.net/problem=46

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

                9  = 7  + 2×12
                15 = 7  + 2×22
                21 = 3  + 2×32
                25 = 7  + 2×32
                27 = 19 + 2×22
                33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
'''

def primeSieve(n):
    sieve = [True] * n
    for i in range(3, int(n**0.5) + 1, 2):
        if sieve[i]:
            sieve[i*i : : 2*i] = [False] * ((n - i*i - 1) // (2 * i) + 1)
    return [2] + [i for i in range(3, n, 2) if sieve[i]]




def main():
    sqrs   = [2 * (i**2) for i in range(1, 1000)]
    primes = primeSieve(1000)
    odds   = list(filter(lambda n: n not in primes, range(1, 1000, 2)))

    print(odds[:20])

    for i in odds:
        for j in sqrs:
            if i <= j: break
            if i - j in primes: return i, j

if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")

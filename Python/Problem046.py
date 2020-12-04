from time import time


''' 
Problem 46: https://projecteuler.net/problem=46

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

                9 = 7 + 2×12
                15 = 7 + 2×22
                21 = 3 + 2×32
                25 = 7 + 2×32
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

def squares(n):
    i, res = 1, []
    while i**i < n:
        res.append(i**i)
    return res



def main(): # takes too long needs work
    
    primes = primeSieve(100)
    oddNums = [i for i in range(1, 100, 2)]
    
    for i in oddNums:
        if i in primes:
            continue
        
        sqrs = squares(i)[::-1]
        
        for j in sqrs:
            if 2*(i - j)+1 not in primes:
                return i
    
            
if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
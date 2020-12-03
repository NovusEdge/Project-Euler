from time import time

''' 
Problem 14: https://projecteuler.net/problem=14
The following iterative sequence is defined for the set of positive integers:

            n → n/2 (n is even)
            n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
			13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

* NOTE: Once the chain starts the terms are allowed to go above one million.
'''

process = lambda n: n // 2 if n%2 == 0 else 3*n + 1

def collatz(n):
    s=[ n, ]
    while n > 1:
        n = process(n)
        s.append(n)
    s.append(1)
    return s

start = time()
res = [len(collatz(i)) for i in range(10**6)]
ans = res.index(max(res))       

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")



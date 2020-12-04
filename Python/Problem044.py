from time import time
from itertools import permutations as perm

''' 
Problem 44: https://projecteuler.net/problem=44

Pentagonal numbers are generated by the formula, P(n)=n(3n−1)/2. The first ten pentagonal numbers are:

                            1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, P(j) and P(k), for which their sum and difference are pentagonal and D = |P(k) − P(j)| is minimised; what is the value of D?
'''

P = lambda n: n*(3*n - 1)//2
D = lambda i: abs( P(i[0]) - P(i[1]) )

start = time()
ans = 0

D_set = dict()
p_Nums = list(map(P, range(1, 100)))
ans = min(set(map(D, perm(range(1, 1000), 2))))

# needs work ... 

if __name__=="__main__":    
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
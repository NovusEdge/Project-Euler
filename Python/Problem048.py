from time import time

''' 
Problem 48: https://projecteuler.net/problem=48

The series, 1**1 + 2**2 + 3**3 + ... + 10**10 = 10405071317.

Find the last ten digits of the series, 1**1 + 2**2 + 3**3 + ... + 1000**1000.
'''

start = time()
ans = sum([i**i for i in range(1, 1000)])%(10**10)

if __name__=="__main__":    
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

from time import time

''' 
Problem 4: https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
'''

isPall = lambda s : str(s)==str(s)[::-1]
palls = []
start = time()

for i in range(999, 100, -1):
    for j in range(999, 100, -1):
        if isPall(i*j):
            palls.append(i*j)
            
ans = max(palls)

if __name__ == "__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time taken: { time() - start }\n")

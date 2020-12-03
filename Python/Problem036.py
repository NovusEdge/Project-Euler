from time import time

''' 
Problem 36: https://projecteuler.net/problem=36

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
'''

pall = lambda n : str(n) == str(n)[::-1]
check = lambda n : pall(n) and pall(bin(n).replace('0b', ''))
start = time()
ans = sum([i for i in range(1,10**6) if check(i)])

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")


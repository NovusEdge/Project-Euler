from math import factorial
from time import time

''' 
Problem 53: https://projecteuler.net/problem=53

There are exactly ten ways of selecting three from five, 12345:
        123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

How many, not necessarily distinct, values of C(n, r) for, 1 ≤ n ≤ 100 are greater than one-million?
'''

C = lambda n, r: factorial(n) // (factorial(n-r) * factorial(r))

def main():
    counter = 0
    for a in range(0, 101):
        for r in range(0, a+1):
            if C(a, r) > pow(10, 6):
                counter += 1
    return counter

if __name__=="__main__":    
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
    

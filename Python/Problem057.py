from math import log10
from time import time

''' 
Problem 57: https://projecteuler.net/problem=57

It is possible to show that the square root of two can be expressed as an infinite continued fraction.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than the denominator?
'''

def main():
    Ln, n, a, b = 1000, 3, 2, 0
    for i in range(2, Ln+1):
        n, a = n + 2*a, n + a
        if int(log10(n)) > int(log10(a)): 
            b += 1
    return b

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")

      

from itertools import islice,count
from math import floor,sqrt
from time import time
from collections import deque

''' 
Problem 35: https://projecteuler.net/problem=35

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
'''

def prime(n):
    if n == 0 or n == 1:
        return None
    if n == 2:
        return True
    return all([not(n % i == 0) for i in islice(count(2),floor(sqrt(n)))])

def checkCircular(n):
    A = deque(list(str(n)))
    
    if not prime(n):
        return False
    
    SomeLis = []
    
    for i in range(len(A)):
        SomeLis.append(int(''.join(list(A))))
        A.rotate()
        
    return all(prime(i) for i in SomeLis)

def main():
    counter = 0
    k = []
    
    for i in range(2, 10**6):
        if checkCircular(i):
            counter += 1
            
    return counter

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
 
    

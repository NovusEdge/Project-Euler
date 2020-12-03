from math import sqrt; from itertools import count, islice, permutations
from time import time
  
''' 
Problem 41: https://projecteuler.net/problem=41

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
'''  
  
prime = lambda n: n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))
isPan = lambda n: sorted(str(n)) == list(str(n))
check = lambda n: prime(n) and isPan(n)

def main():
    perms = permutations("1234567", 7)
    checkSet = set()
    for i in perms:
        temp = int(''.join(i))
        if prime(temp):
            checkSet.add(temp)
        
    return max(checkSet)

if __name__=="__main__":
    start = time() 
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")

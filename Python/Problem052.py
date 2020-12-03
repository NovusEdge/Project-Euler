from time import time

''' 
Problem 52: https://projecteuler.net/problem=52

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
'''

sameDigits = lambda a, b: sorted(str(a)) == sorted(str(b))
i = 10**5
start = time()

while not (sameDigits(i,2*i) and sameDigits(i,3*i) and sameDigits(i,4*i) and sameDigits(i,5*i) and sameDigits(i,6*i)): i += 1

if __name__=="__main__":    
    print(f"\nAnswer: { i }")
    print(f"Time Taken: { time() - start }\n")
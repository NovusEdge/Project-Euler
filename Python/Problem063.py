from time import time

''' 
Problem 63: https://projecteuler.net/problem=63


The 5-digit number, 16807=7**5, is also a fifth power. Similarly, the 9-digit number, 134217728=8**9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?
'''

counter = 0

for i in range(100):
    for j in range(100):
        if len(str(i**j)) == j:
            counter += 1

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { counter - 1 }") # the -1 is for "1" 
    print(f"Time Taken: { time() - start }\n")
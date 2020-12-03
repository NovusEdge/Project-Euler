from time import time

''' 
Problem 16: https://projecteuler.net/problem=16

What is the sum of the digits of the number 2**1000?
'''

start = time()
ans = sum(map(int, [i for i in str(2**1000)]))

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
from time import time

''' 
Problem 56: https://projecteuler.net/problem=56
Considering natural numbers of the form, a**b, where a, b < 100, what is the maximum digital sum?
'''

start = time()
digitSum = lambda n: sum(map(int, str(n)))
refList = []

for a in range(99, 1, -1):
    for b in range(99, 1, -1):
        refList.append(digitSum(a**b))

ans = max(refList)

if __name__=="__main__":
   print(f"\nAnswer: { ans }")
   print(f"Time Taken: { time() - start }\n")


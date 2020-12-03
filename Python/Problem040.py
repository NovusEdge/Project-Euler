from time import time

''' 
Problem 40: https://projecteuler.net/problem=40

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.
If dn represents the nth digit of the fractional part, find the value of the following expression:

                d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
'''

start = time()
chamConst = ''.join(map(str, range(1000001)))
d = lambda n: int(chamConst[n])
ans = d(1)*d(10)*d(100)*d(1000)*d(10000)*d(100000)*d(1000000)

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
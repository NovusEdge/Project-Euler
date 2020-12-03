import num2words
from time import time

''' 
Problem 17: https://projecteuler.net/problem=17

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

* NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
'''

start = time.time()
word = lambda k : num2words.num2words(k).replace(' ', '').replace('-', '')
p = ''

for i in range(1, 1001):
    p += word(i)
    
ans = len(p)

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

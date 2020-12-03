from math import factorial 
from time import time

''' 
Problem 74: https://projecteuler.net/problem=74

The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:

                        1! + 4! + 5! = 1 + 24 + 120 = 145

Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:

                        169 → 363601 → 1454 → 169
                        871 → 45361 → 871
                        872 → 45362 → 872

It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,

                        69 → 363600 → 1454 → 169 → 363601 (→ 1454)
                        78 → 45360 → 871 → 45361 (→ 871)
                        540 → 145 (→ 145)

Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.

How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
'''

digit_fact_sum = lambda n : sum( factorial( int( i ) ) for i in str(n) )


def check_chain(n):
    refLis, i = [n], n
    counter = 0
    if n == digit_fact_sum(n):
        return 1
    while True:
        i = digit_fact_sum(i)
        refLis.append(i)
        counter += 1
        if list(dict.fromkeys(refLis)) != refLis:
            return counter 


def main():
    counter = 0
    for i in range(1, 10**6 + 1):
        if check_chain(i) == 60:
            counter += 1
    return counter


if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")


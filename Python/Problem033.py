from time import time

''' 
Problem 33: https://projecteuler.net/problem=33

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
'''

def cut(a, b):
    a1 = list(str(a))
    b1 = list(str(b))
    for i in a1:
        if i in b1:
            a1.pop(a1.index(i))
            b1.pop(b1.index(i))
    return int(''.join(a1)),int(''.join(b1))


def check(a, b):
    if a%10 == 0 and b%10 == 0 or a == b:
        return False
    return cut(a, b)[0] / cut(a, b)[1] == a/b



def main():
    s=[]
    for a in range(11, 100):
        for b in range(11, 100):
            try:
                if check(a, b):
                    s.append(cut(a, b))
            except ZeroDivisionError:
                pass
    return s[:4]

if __name__=="__main__":
    start = time()
    print(f"\nResultant Fractions: { main() }")
    print(f"Time Taken: { time() - start }\n")

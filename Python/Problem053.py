from math import factorial
from time import time

C = lambda n, r: factorial(n) // (factorial(n-r) * factorial(r))

def main():
    counter = 0
    for a in range(0, 101):
        for r in range(0, a+1):
            if C(a, r) > pow(10, 6):
                counter += 1
    return counter

if __name__=="__main__":    
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
    

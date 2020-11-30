from time import time
from math import factorial

digitFact = lambda n: sum(map(lambda x: factorial(int(x)), str(n)))

ans = 0
for i in range(3, 100000):
    if digitFact(i) == i:
        ans += i

if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

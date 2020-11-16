from math import sqrt
from time import time
from functools import reduce

def factors(n):    
    return set(reduce(list.__add__, 
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))

start = time()
temp_list = list(filter(lambda x: sum(factors(x))-x > x, range(1, 28124)))
boolMap = dict()

for i in range(1, 28124): boolMap[i] = True

for i in temp_list:
    for j in temp_list:
        boolMap[i+j] = False

ans = sum([i for i in boolMap if boolMap[i]])

if __name__ == "__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time taken:{ time() - start }\n")

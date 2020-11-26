from time import time
from itertools import islice, count
from math import sqrt
from functools import reduce

prime = lambda n: n > 1 and all(n % i for i in islice(count(2), int(sqrt(n) - 1)))
factors = lambda n: set(reduce(list.__add__, ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))
digits = lambda n: list(map(int, str(n)))
def check(n):
    flag_2 = int(str(n)[-1])%2==0
    flag_3 = sum(digits)%3==0
    flag_4 = int(str(n)[-2:-1])%4==0
    flag_5 = str(n)[-1] in ['5', '0']
    flag_6 = flag_2 and flag_3
    ...
    


start = time()
ans = 2520


if __name__ == '__main__':
    print(f"Answer: { ans }")
    print(f"Time Taken: { time() - start }")
    
    

    
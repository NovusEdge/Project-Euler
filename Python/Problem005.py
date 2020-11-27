from time import time
from functools import reduce
from itertools import starmap

check = lambda n: all(map(lambda l: l[1]==0, starmap(int.__divmod__, zip([n]*20, range(2, 21)))))
start = time() 
ans = 2520
while not check(ans): ans += 10

if __name__ == '__main__':
    print(f"Answer: { ans }")
    print(f"Time Taken: { time() - start }")
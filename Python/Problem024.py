from time import time
from itertools import permutations as perm

start = time()
perms = list(map(lambda x: ''.join(x), perm("0123456789", 10)))
ans = sorted(perms)[999999]

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
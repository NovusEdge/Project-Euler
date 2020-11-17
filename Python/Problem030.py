from time import time

digit = lambda n: map(int, str(n))   
check = lambda n: sum([ i**5 for i in digit(n) ]) == n


start = time()
ans = sum(filter(check, range(10**3, 10**6)))

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
from time import time

start = time()
ans = sum([i for i in range(1001) if (i % 3 == 0 or i % 5 == 0)])

if __name__ == '__main__':
    print(f"\nAnswer:{ ans }")
    print(f"Time Taken:{ time() - start }\n")

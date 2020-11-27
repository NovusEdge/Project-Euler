from time import time

isPall = lambda s : str(s)==str(s)[::-1]
palls = []
start = time()

for i in range(999, 100, -1):
    for j in range(999, 100, -1):
        if isPall(i*j):
            palls.append(i*j)
            
ans = max(palls)

if __name__ == "__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time taken: { time() - start }\n")

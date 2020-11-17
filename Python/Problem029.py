from time import time
start = time()
powers = []

for a in range(2, 101):
    for b in range(2, 101):
        powers.append(a ** b)

ans = len(set(powers))
print(100**100)
if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")

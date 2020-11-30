from time import time

start = time()
ans = ( (28433 * 2**7830457) + 1 ) % (10**10)

if __name__=="__main__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
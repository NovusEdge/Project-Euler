from time import time

sameDigits = lambda a, b: sorted(str(a)) == sorted(str(b))
i = 10**5
start = time()

while not (sameDigits(i,2*i) and sameDigits(i,3*i) and sameDigits(i,4*i) and sameDigits(i,5*i) and sameDigits(i,6*i)): i += 1

if __name__=="__main__":    
    print(f"\nAnswer: { i }")
    print(f"Time Taken: { time() - start }\n")
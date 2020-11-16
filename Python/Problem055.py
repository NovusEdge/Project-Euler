from time import time


def isLychrel(n):
	for i in range(50):
		n += int(str(n)[:: -1])
		if str(n) == str(n)[:: -1]:
			return False
	return True



start = time()
ans = len(list(filter(isLychrel, range(10000))))

if __name__=="__main__":
   print(f"\nAnswer: { ans }")
   print(f"Time Taken: { time() - start }\n")

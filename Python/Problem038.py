from time import time

''' 
Problem 38: https://projecteuler.net/problem=38

Take the number 192 and multiply it by each of 1, 2, and 3:

        192 × 1 = 192
        192 × 2 = 384
        192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
'''

def checkPan(n):
    return list("123456789") == sorted(str(n))

def getProduct(n):
    i = 1
    
    while True:
        k = range(1, i+1)
        tempVar = ''.join([str(n*i) for i in k])
        if len(tempVar) > 9:
            return 0
        if len(tempVar) == 9:
            if checkPan(int(tempVar)):
                return int(tempVar)
        i += 1


def main():
    maxProduct = 1
    for i in range(1,10000):
        if maxProduct < getProduct(i):
            maxProduct = getProduct(i)
    return maxProduct 


if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
    

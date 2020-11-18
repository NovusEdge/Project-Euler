from time import time

def checkPan(n):
    '''To check if a number is a 9-digit pandigital number'''
    return list("123456789") == sorted(str(n))

def getProduct(n):
    
    '''
    This function is to get the concatenation of the products of number and list of numbers till it hits 9 digits or more
    '''
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
    '''Main Function to loop through numbers and runs the getProduct function on each'''
    maxProduct = 1
    for i in range(1,10000):
        if maxProduct < getProduct(i):
            maxProduct = getProduct(i)
    return maxProduct 


if __name__=="__main__":
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
    

from time import time

''' 
Problem 42: https://projecteuler.net/problem=42

The nth term of the sequence of triangle numbers is given by, t(n) = n(n+1)/2; so the first ten triangle numbers are:

                                1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt, a 16K text file containing nearly two-thousand common English words, how many are triangle words?
'''

T = lambda n : n * (n + 1) // 2
tNums = [T(i) for i in range(1000)]
nameNum = lambda s : sum([ord(i) - 64 for i in s])
check = lambda s : nameNum(s) in tNums

def getWords():
    tempStr = ''    
    with open('p042_words.txt', 'r') as f:
        tempStr = ''.join(f.readlines())
    out = tempStr.split(',')
    return sorted([i[1:len(i)-1] for i in out])

start = time() 
ans = sum([1 for i in getWords() if check(i)])

if __name__=="__main__":    
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")



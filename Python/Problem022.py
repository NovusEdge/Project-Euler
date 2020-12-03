from time import time

''' 
Problem 22: https://projecteuler.net/problem=22

Using names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
'''

nameNum = lambda s : sum([ord(i) - 64 for i in s])

def getWords():
    tempStr = ''    
    with open('./p022_names.txt', 'r') as f:
        for i in f.readlines():
            tempStr += i
        out = tempStr.split(',')
    k = []
    for i in out:
        k.append(i[ 1: len(i) - 1 ])
    return sorted(k)


def main():
    total = 0
    words = getWords()
    print(words)
    for i in range(len(words)):
        total += nameNum(words[i]) * (i + 1)
    return total

if __name__ == "__main__":
     start = time()
     print(f"\nSum of all namescores in file:{ main() }")
     print(f"Time taken:{ time() - start }\n")
 

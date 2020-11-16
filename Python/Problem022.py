from time import time

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
 

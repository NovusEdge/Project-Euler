from time import time
isPan = lambda n: sorted(n) == list("123456789")


def main():
    resArr = set()
    for i in range(1, 2000):
        for j in range(1, 2000):
            if isPan(f"{ i }{ j }{ i*j }"):
                resArr.add(i * j)
    return sum(resArr)


if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
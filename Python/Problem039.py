from time import time

isIntTriangle = lambda a,b,c: c**2 == a**2 + b**2 and a < b < c
sols, pFlag = 0, 0

def p_sol(p):
    sol_counter = 0
    for i in range(400):
        for j in range(400):
            k = p - i - j
            if isIntTriangle(i, j, k) and k > 0:
                sol_counter += 1
    return sol_counter

start = time()

for i in range(1001):
    temp = p_sol(i)
    if temp > sols:
        print(i)
        sols = temp
        pFlag = i


if __name__ == '__main__':
    print(f"Answer: { pFlag }")
    print(f"Time Taken: { time() - start }") 

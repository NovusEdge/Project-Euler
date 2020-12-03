from time import time

''' 
Problem 39: https://projecteuler.net/problem=39

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

        {20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
'''

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
    print(f"\nAnswer: { pFlag }")
    print(f"Time Taken: { time() - start }\n") 

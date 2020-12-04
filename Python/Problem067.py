from time import time

''' 
Problem 67: https://projecteuler.net/problem=67

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

                                                3
                                               7 4
                                              2 4 6
                                             8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt, a 15K text file containing a triangle with one-hundred rows.
'''

triangle = open('./p067_triangle.txt', 'r').read()

# there might be a newline in the file thats at the end of the file, be sure to remove it first...

start = time()

triangle = list(map(str.split, triangle.split('\n')))
triangle = [list(map(int, i)) for i in triangle]

for i in range(len( triangle )-1)[::-1]:
    for j in range(len(triangle[i])):
        triangle[i][j] += max( triangle[i + 1][j], triangle[i + 1][j + 1] )
        
ans = triangle[0][0]

if __name__ == '__main__':
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
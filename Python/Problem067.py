from time import time

triangle = open('./p067_triangle.txt', 'r').read()

# there might be a newline in the file thats at the end of the file, be sure to remove it first...

start = time()

triangle = list(map(lambda x: x.split(' '), triangle.split('\n')))
triangle = [list(map(int, i)) for i in triangle]

for i in range(len( triangle )-1)[::-1]:
    for j in range(len(triangle[i])):
        triangle[i][j] += max( triangle[i + 1][j], triangle[i + 1][j + 1] )
        
ans = triangle[0][0]

if __name__ == '__main__':
    print(f"\nAnswer:{ ans }")
    print(f"Time Taken:{ time() - start }\n")
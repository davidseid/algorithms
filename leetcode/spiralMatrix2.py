
def generateMatrix(n):
    matrix = [[0 for r in range(0, n)] for c in range(0, n)] 

    num = 1
    off = 0

    while num <= n * n:

        for i in range(off, n - off): 
            matrix[off][i] = num
            num += 1
        
        for i in range(off + 1, n - off):
            matrix[i][n - 1 - off] = num
            num += 1
        
        for i in range(n - 2 - off, off - 1, -1):
            matrix[n - 1 - off][i] = num
            num += 1

        for i in range(n - 2 - off, off, -1):
            matrix[i][off] = num
            num += 1
        
        off += 1


    return matrix

print(generateMatrix(5))
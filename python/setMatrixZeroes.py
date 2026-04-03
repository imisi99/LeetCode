from re import T


def setZeroes(matrix: list[list[int]]) -> None:
    m = 0
    zeroes = []
    while m < len(matrix):
        n = 0
        while n < len(matrix[0]):
            if matrix[m][n] == 0:
                zeroes.append([m, n])
            n += 1
        m += 1

    for zero in zeroes:
        rowCol(matrix, zero[0], zero[1])


def rowCol(matrix: list[list[int]], m: int, n: int):
    start, end = 0, len(matrix)
    while start < end:
        matrix[start][n] = 0
        start += 1

    start, end = 0, len(matrix[0])
    while start < end:
        matrix[m][start] = 0
        start += 1


# space -> 0(1)
def setZeroesI(matrix: list[list[int]]) -> None:
    zeroRow = False
    zeroCol = False

    row, col = len(matrix), len(matrix[0])
    m, n = 0, 0
    while m < row:
        if matrix[m][0] == 0:
            zeroCol = True
            break
        m += 1

    while n < col:
        if matrix[0][n] == 0:
            zeroRow = True
            break
        n += 1

    m = 1
    while m < row:
        n = 1
        while n < col:
            if matrix[m][n] == 0:
                matrix[m][0] = 0
                matrix[0][n] = 0
            n += 1
        m += 1

    m = 1
    while m < row:
        if matrix[m][0] == 0:
            n = 1
            while n < col:
                matrix[m][n] = 0
                n += 1
        m += 1

    n = 1
    while n < col:
        if matrix[0][n] == 0:
            m = 1
            while m < row:
                matrix[m][n] = 0
                m += 1
        n += 1

    if zeroRow:
        n = 0
        while n < col:
            matrix[0][n] = 0
            n += 1

    if zeroCol:
        m = 0
        while m < row:
            matrix[m][0] = 0
            m += 1


matrix = [
    [1, 1, 1],
    [1, 0, 1],
    [1, 1, 1],
]

matrix1 = [
    [0, 1, 0],
    [1, 1, 1],
    [1, 1, 1],
]

setZeroes(matrix)
setZeroesI(matrix1)
print(matrix)
print(matrix1)

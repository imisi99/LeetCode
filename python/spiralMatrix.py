def spiralMatrix(matrix: list[list[int]]) -> list[int]:
    array = []
    spiral(0, len(matrix[0]) - 1, 0, len(matrix) - 1, matrix, array)
    return array


# Time -> 0(M*N) Space ->  0(M*N)
def spiral(
    startCol, endCol, startRow, endRow, matrix: list[list[int]], array: list[int]
):
    while startCol <= endCol and startRow <= endRow:
        # go right
        right = startCol
        while right <= endCol:
            array.append(matrix[startRow][right])
            right += 1

        startRow += 1

        if startRow > endRow:
            break

        # go down
        down = startRow
        while down <= endRow:
            array.append(matrix[down][endCol])
            down += 1

        endCol -= 1

        if startCol > endCol:
            break

        # go left
        left = endCol
        while left >= startCol:
            array.append(matrix[endRow][left])
            left -= 1

        endRow -= 1
        if startRow > endRow:
            break

        # go up
        up = endRow
        while up >= startRow:
            array.append(matrix[up][startCol])
            up -= 1

        startCol += 1


print(
    spiralMatrix(
        [
            [1, 2, 3, 4],
            [5, 6, 7, 8],
            [9, 10, 11, 12],
            [13, 14, 15, 16],
            [17, 18, 19, 20],
            [21, 22, 23, 24],
        ]
    )
)

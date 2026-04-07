# Time -> 0(log(M*N))
def searchMatrix(matrix: list[list[int]], target: int) -> bool:
    if len(matrix) == 0 or len(matrix[0]) == 0:
        return False
    start, end = 0, len(matrix) - 1
    bestMid = 0
    while start <= end:
        mid = (start + end) // 2

        val = matrix[mid][0]
        if val == target:
            return True
        elif val < target:
            bestMid = mid
            start = mid + 1
        else:
            end = mid - 1

    start, end = 0, len(matrix[0]) - 1
    while start <= end:
        mid = (start + end) // 2

        val = matrix[bestMid][mid]
        if val == target:
            return True
        elif val < target:
            start = mid + 1
        else:
            end = mid - 1
    return False


matrix = [[]]
print(searchMatrix(matrix, 8))

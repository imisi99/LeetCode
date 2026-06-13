# Time -> 0(N) Space -> 0(N)
def minimumTotal(triangle: list[list[int]]) -> int:
    memo = {}

    def recurse(row: int, col: int) -> int:
        if row >= len(triangle) or col >= len(triangle[row]):
            return 0

        if (row, col) in memo:
            return memo[(row, col)]

        samecol = recurse(row + 1, col)
        diffcol = recurse(row + 1, col + 1)

        memo[(row, col)] = triangle[row][col] + min(samecol, diffcol)
        return memo[(row, col)]

    return recurse(0, 0)


print(minimumTotal([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]))

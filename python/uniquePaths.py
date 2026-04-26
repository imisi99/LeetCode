def uniquePaths(m: int, n: int) -> int:
    memo = [[0 for _ in range(n)] for _ in range(m)]
    return searchPaths(0, 0, m, n, memo)


# Time -> 0(M*N) Space -> 0(M*N)
def searchPaths(i: int, j: int, m: int, n: int, memo: list[list[int]]) -> int:
    if i == m or j == n:
        return 0

    if i == m - 1 and j == n - 1:
        return 1

    if memo[i][j] != 0:
        return memo[i][j]

    path = 0

    path += searchPaths(i + 1, j, m, n, memo)
    path += searchPaths(i, j + 1, m, n, memo)

    memo[i][j] = path

    return memo[i][j]


print(uniquePaths(3, 7))

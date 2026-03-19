# Time -> 0(max(N, M)) Space -> 0(1)
def numberOfSubmatrices(grid: list[list[str]]) -> int:
    count = 0
    x_count = 0
    y_count = 0

    for val in grid[0]:
        if val == "X":
            x_count += 1
        elif val == "Y":
            y_count += 1

        if x_count == y_count and x_count != 0:
            count += 1

    x_count = 0
    y_count = 0
    for idx in range(len(grid)):
        val = grid[0][idx]
        if val == "X":
            x_count += 1
        elif val == "Y":
            y_count += 1

        if x_count == y_count and x_count != 0:
            count += 1

    return count


print(numberOfSubmatrices([["X", "Y", "."], ["Y", ".", "."]]))
print(numberOfSubmatrices([["X", "X"], ["X", "Y"]]))
print(numberOfSubmatrices([[".", "."], [".", "."]]))

def numTrees(n: int) -> int:
    sum = [0]
    recurse(1, n, sum)
    return sum[0]


def recurse(idx: int, n: int, sum: list[int]):
    if idx == n:
        sum[0] += 1
        return

    if idx > n:
        return

    recurse(idx + 1, n, sum)
    recurse(idx + 1, n, sum)
    recurse(idx + 2, n, sum)


print(numTrees(4))

def climbStairs(n: int) -> int:
    memo = {}
    return countStairs(0, n, memo)


# Time -> 0(N) Space -> 0(N)
def countStairs(idx: int, n: int, memo: dict[int, int]) -> int:
    if idx >= n:
        if idx == n:
            return 1
        return 0

    if idx in memo:
        return memo[idx]

    oneStep = countStairs(idx + 1, n, memo)
    twoStep = countStairs(idx + 2, n, memo)

    memo[idx] = oneStep + twoStep

    return memo[idx]

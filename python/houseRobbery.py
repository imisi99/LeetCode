def houseRobbery(profits: list[int]) -> int:
    memo = [-1] * len(profits)
    return max(findMax(0, memo, profits), findMax(1, memo, profits))


# Time -> 0(N) Space -> 0(N)
def findMax(i: int, memo: list[int], profits: list[int]) -> int:
    if i >= len(profits):
        return 0

    if memo[i] != -1:
        return memo[i]
    memo[i] = profits[i] + max(
        findMax(i + 2, memo, profits), findMax(i + 3, memo, profits)
    )
    return memo[i]


print(houseRobbery([1, 2, 3, 1]))

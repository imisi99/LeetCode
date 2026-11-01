# Time -> 0(N*N*N) Space -> 0(1)
def minMoves(nums: list[int]) -> int:
    count = 0
    minVal, maxVal = minMax(nums)
    while minVal != maxVal:
        skip = True
        for i in range(len(nums)):
            if (nums[i] == maxVal) and skip:
                skip = False
                continue
            nums[i] += 1
        count += 1
        minVal, maxVal = minMax(nums)
    return count


def minMax(nums: list[int]):
    minVal = float("inf")
    maxVal = float("-inf")

    for val in nums:
        if val < minVal:
            minVal = val
        if val > maxVal:
            maxVal = val

    return minVal, maxVal


# Time -> 0(N*N) Space -> 0(1)
def minMovesII(nums: list[int]) -> int:
    nums.sort()
    count = 0

    for i in range(1, len(nums)):
        diff = nums[i] - nums[0]
        if diff > 0:
            for j in range(len(nums)):
                if i == j:
                    continue
                nums[j] += diff
            count += diff

    return count


print(minMoves([1, 2, 3]))
print(minMovesII([1, 2, 3]))
print(minMovesII([1, 10000000]))
print(minMoves([1, 10000000]))

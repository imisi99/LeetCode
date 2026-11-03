# Time -> 0(N) Space -> 0(1)
def minMoves(nums: list[int]) -> int:
    count = 0
    maxVal = 0

    for num in nums:
        if num > maxVal:
            maxVal = num

    for num in nums:
        count += maxVal - num

    return count


print(minMoves([1, 2, 3]))
print(minMoves([4, 4, 5]))

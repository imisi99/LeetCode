# Time -> 0(Nd) Space -> 0(d)
def smallestIndex(nums: list[int]) -> int:
    for idx, val in enumerate(nums):
        currSum = 0
        for v in str(val):
            currSum += int(v)
        if idx == currSum:
            return idx

    return -1


print(smallestIndex([1, 3, 2]))
print(smallestIndex([1, 10, 11]))
print(smallestIndex([1, 2, 3]))

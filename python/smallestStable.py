# Time -> 0(N) Space -> 0(N)
def firstStableIndex(nums: list[int], k: int) -> int:
    lookup = {}

    maxVal = nums[0]
    for idx in range(len(nums)):
        if nums[idx] > maxVal:
            maxVal = nums[idx]
        lookup[idx] = [maxVal]

    idx = len(nums) - 1
    minVal = nums[idx]
    for idx in reversed(range(len(nums))):
        if nums[idx] < minVal:
            minVal = nums[idx]
        lookup[idx].append(minVal)

    for idx in range(len(nums)):
        score = lookup[idx]
        stable = score[0] - score[1]
        if stable <= k:
            return idx
    return -1


print(firstStableIndex([5, 0, 1, 4], 3))
print(firstStableIndex([3, 2, 1], 1))
print(firstStableIndex([0], 0))
print(firstStableIndex([6, 1, 4], 5))

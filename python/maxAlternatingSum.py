# Time => 0(NlogN) Space => 0(1)
def maxAlternatingSum(nums: list[int]) -> int:
    for i in range(len(nums)):
        nums[i] = nums[i] * nums[i]

    nums.sort()

    k = len(nums) // 2

    subtract = nums[:k]
    add = nums[k:]

    total = 0
    for val in add:
        total += val

    for val in subtract:
        total -= val

    return total


print(maxAlternatingSum([1, 2, 3]))
print(maxAlternatingSum([1, -1, 2, -2, 3, -3]))

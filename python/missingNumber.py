# Time -> 0(N) Space -> 0(1)
def missingNumber(nums: list[int]) -> int:
    expectedSum = 0
    actualSum = 0

    for i in range(len(nums) + 1):
        expectedSum += i
        if i != len(nums):
            actualSum += nums[i]

    return expectedSum - actualSum


print(missingNumber([1, 2]))
print(missingNumber([1, 3, 0]))

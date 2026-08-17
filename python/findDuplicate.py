# Time -> 0(N) Space -> 0(1)
def findDuplicate(nums: list[int]) -> int:
    actualSum = 0
    expSum = 0

    i = 1
    while i < len(nums):
        expSum += i
        i += 1

    for val in nums:
        actualSum += val

    return actualSum - expSum


print(findDuplicate([1, 2, 3, 4, 5, 6, 6]))

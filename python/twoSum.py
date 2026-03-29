# Time -> 0(N) Space -> 0(N)
def twoSum(nums: list[int], target: int) -> list[int]:
    seen = {}
    for i in range(len(nums)):
        seen[nums[i]] = i

    for i in range(len(nums)):
        diff = target - nums[i]
        if diff in seen and i != seen[diff]:
            return [i, seen[diff]]

    return []


print(twoSum([2, 7, 11, 15], 9))

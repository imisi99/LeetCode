# Time -> 0(N) Space -> 0(N)
def isGood(nums: list[int]) -> bool:
    base = len(nums) - 1
    baseMap = [0] * (base + 1)
    for num in nums:
        if num <= 0 or num > base:
            return False
        baseMap[num] += 1

        if baseMap[num] > 2 or (baseMap[num] == 2 and num != base):
            return False

    return True


print(isGood([1, 2, 3, 3]))
print(isGood([0, 1, 2, 3]))

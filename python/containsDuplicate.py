# Time -> 0(N) Space -> 0(N)
def contains(nums: list[int]) -> bool:
    duplicate = {}
    for num in nums:
        exists = duplicate.get(num, False)
        if exists:
            return True
        duplicate[num] = True
    return False


print(contains([0, 2, 3, 4, 4]))
print(contains([0, 2, 3, 4, 5]))

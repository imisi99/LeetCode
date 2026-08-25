# Time -> 0(N) Space -> 0(N)
def missingMultiple(nums: list[int], k: int) -> int:
    lookup = set()
    for num in nums:
        lookup.add(num)

    mul = 1
    while 1 <= len(nums):
        if mul * k not in lookup:
            return mul * k
        mul += 1

    return (mul + 1) * k


# Time -> 0(NlogN) Space -> 0(1)
def missingMultipleI(nums: list[int], k: int) -> int:
    nums.sort()

    mul = k
    for num in nums:
        if num > mul:
            return mul
        if num == mul:
            mul += k

    return mul


print(missingMultiple([2, 3, 6], 2))
print(missingMultipleI([2, 3, 6], 2))

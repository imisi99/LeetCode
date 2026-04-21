# Time -> 0(N) Space -> 0(1)
def singleNumber(nums: list[int]) -> int:
    single = 0

    for val in nums:
        single ^= val
    return single


print(singleNumber([1, 1, 2, 3, 4, 3, 2]))

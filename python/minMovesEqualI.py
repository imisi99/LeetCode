# Time -> 0(NlogN) Space -> 0(1)
def minMoves(nums: list[int]) -> int:
    nums.sort()
    mid = 0
    if len(nums) % 2 == 1:
        mid = nums[len(nums) // 2]
    else:
        mid = (nums[len(nums) // 2] + nums[((len(nums) // 2) - 1)]) // 2

    count = 0
    for val in nums:
        num = val - mid
        num = abs(num)
        count += num
    return count


print(minMoves([1, 2, 3]))

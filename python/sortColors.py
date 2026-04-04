# Time -> 0(N) Space -> 0(1)
def sortColors(nums: list[int]):
    red, blue = 0, 2
    redIdx, blueIdx = 0, len(nums) - 1
    idx = 0

    while idx <= blueIdx:
        val = nums[idx]
        if val == red:
            nums[idx], nums[redIdx] = nums[redIdx], nums[idx]
            redIdx += 1
            idx += 1
        elif val == blue:
            nums[idx], nums[blueIdx] = nums[blueIdx], nums[idx]
            blueIdx -= 1
        else:
            idx += 1


array = [2, 0, 2, 0, 2, 1, 1, 0]
print(array)
sortColors(array)
print(array)

# Time -> 0(N) Space -> 0(1)
def removeDuplicates(nums: list[int]) -> int:
    idxToSwap = 0
    idx = 0
    while idx < len(nums):
        currIdx = idx
        while idx < len(nums) and nums[idx] == nums[currIdx]:
            idx += 1

        if idx - currIdx - 1 >= 1:
            nums[idxToSwap] = nums[currIdx]
            nums[idxToSwap + 1] = nums[currIdx]
            idxToSwap += 2
        else:
            nums[idxToSwap] = nums[currIdx]
            idxToSwap += 1
    return idxToSwap


array = [0, 0, 1, 1, 1, 1, 2, 3, 3]
print(array)
print(removeDuplicates(array))
print(array)

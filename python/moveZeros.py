# Time -> 0(N) Space -> 0(1)
def moveZeros(nums: list[int]):
    idxToSwap, idx = 0, 0
    while idx < len(nums):
        if nums[idx] != 0:
            nums[idxToSwap], nums[idx] = nums[idx], nums[idxToSwap]
            idxToSwap += 1

        idx += 1


array = [4, 2, 4, 0, 0, 3, 0, 5, 1, 0]
print(array)
moveZeros(array)
print(array)
